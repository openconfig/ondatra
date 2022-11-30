// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ixweb

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"golang.org/x/net/context"

	closer "github.com/openconfig/gocloser"
)

const (
	viewReadyAttempts = 30

	// EgressStatsCaption is the name of the egress statistics view.
	EgressStatsCaption = "EgressStatView"
	// TrafficItemStatsCaption is the name of the trafic item statics view.
	TrafficItemStatsCaption = "Traffic Item Statistics"

	viewsPath   = "statistics/view"
	viewOpsPath = viewsPath + "/operations"
)

var (
	csvCounter  uint32
	drilldownFn = (*StatView).drilldown
)

// Stats represents the statistics API for an IxNetwork session.
type Stats struct {
	sess *Session
}

// Views fetches the statistics views for the session and returns them as a
// map from the view caption to a representation of the view.
func (s *Stats) Views(ctx context.Context) (map[string]*StatView, error) {
	out := []struct {
		ID      int    `json:"id"`
		Caption string `json:"caption"`
	}{}
	if err := s.sess.Get(ctx, "statistics/view", &out); err != nil {
		return nil, fmt.Errorf("error fetching statistics views: %w", err)
	}
	views := make(map[string]*StatView)
	for _, v := range out {
		views[v.Caption] = &StatView{sess: s.sess, id: v.ID, caption: v.Caption}
	}
	return views, nil
}

// ConfigEgressView will create a statistics views for the specified traffic
// item names, broken down by their egress-tracked values. If an egress stat
// view already exists, it will be deleted first before the new one is created.
// The egressPageSize is the number of unique values that will be tracked.
func (s *Stats) ConfigEgressView(ctx context.Context, trafficItems []string, egressPageSize int) (*StatView, error) {
	const minEgressPageSize, maxEgressPageSize = 1, 79
	if egressPageSize < minEgressPageSize || egressPageSize > maxEgressPageSize {
		return nil, fmt.Errorf("egress page size %v not in allowed range [%v, %v]", egressPageSize, minEgressPageSize, maxEgressPageSize)
	}
	views, err := s.Views(ctx)
	if err != nil {
		return nil, err
	}
	if egressView, ok := views[EgressStatsCaption]; ok {
		if err := s.sess.Delete(ctx, egressView.path()); err != nil {
			return nil, err
		}
	}
	createView := struct {
		Caption          string `json:"caption"`
		TreeViewNodeName string `json:"treeViewNodeName"`
		Type             string `json:"type"`
		Visible          bool   `json:"visible"`
	}{
		Caption:          EgressStatsCaption,
		TreeViewNodeName: "Egress Custom Views",
		Type:             "layer23TrafficFlow",
		Visible:          true,
	}
	egressLinks := struct {
		Links []struct {
			Href string
		}
	}{}
	if err := s.sess.Post(ctx, "statistics/view", createView, &egressLinks); err != nil {
		return nil, fmt.Errorf("error creating Egress Stat View: %w", err)
	}
	if len(egressLinks.Links) != 1 {
		return nil, fmt.Errorf("expected one link to the egress view, got: %v", egressLinks)
	}
	idStr := filepath.Base(egressLinks.Links[0].Href)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("error converting egress view id %q to int: %w", idStr, err)
	}
	egressView := &StatView{sess: s.sess, id: id, caption: EgressStatsCaption}
	if err := egressView.configEgressFilters(ctx, trafficItems); err != nil {
		return nil, err
	}

	// IxNetwork has a maximum number of rows per page of 2000, so the following
	// condition must be true: pageSize * (egressPageSize + 1) <= 2000.
	// In addition, IxNetwork requires the pageSize to be a multiple of 25.
	const (
		maxPageSize      = 2000
		pageSizeMultiple = 25
	)
	pageSize := ((maxPageSize / (egressPageSize + 1)) / pageSizeMultiple) * pageSizeMultiple
	pageSizes := struct {
		PageSize       int `json:"pageSize"`
		EgressPageSize int `json:"egressPageSize"`
	}{
		PageSize:       pageSize,
		EgressPageSize: egressPageSize,
	}
	dataPath := path.Join(egressView.path(), "data")
	if err := egressView.sess.Patch(ctx, dataPath, pageSizes); err != nil {
		return nil, fmt.Errorf("error setting page sizes: %w", err)
	}
	return egressView, nil
}

// StatView is an IxNetwork statistics view.
type StatView struct {
	sess    *Session
	id      int
	caption string
}

func (v *StatView) path() string {
	return fmt.Sprintf("%s/%d", viewsPath, v.id)
}

func (v *StatView) String() string {
	return v.caption + " View"
}

// Caption returns the caption of the view.
func (v *StatView) Caption() string {
	return v.caption
}

// StatTable is a list of StatRows in an IxNetwork statistics view.
type StatTable []StatRow

// StatRow is an individual row of statistics in an IxNetwork statistics view.
// The key of the map is the column name and the value is the statistic value.
// IxNetwork always provides the value of statistics as strings.
type StatRow map[string]string

// FetchTable fetches a table of statistic values from the view.
// If drilldowns are specified, the results of each drilldown on each row of
// the target table are also returned.
func (v *StatView) FetchTable(ctx context.Context, drilldowns ...string) (StatTable, error) {
	if v.caption == EgressStatsCaption {
		if len(drilldowns) != 0 {
			return nil, fmt.Errorf("drilldowns not supported for %q", EgressStatsCaption)
		}
		return v.fetchEgressTableFromPages(ctx)
	}
	tbl, err := v.fetchTableFromCSV(ctx)
	if err != nil {
		return nil, err
	}
	return v.fetchDrilldowns(ctx, tbl, drilldowns)
}

func (v *StatView) fetchDrilldowns(ctx context.Context, tbl StatTable, drilldowns []string) (StatTable, error) {
	if len(drilldowns) == 0 {
		return tbl, nil
	}
	views := map[string]*StatView{v.caption: v}
	for rowIdx, row := range tbl {
		viewInScope := views[v.caption]
		ddRow := rowIdx + 1
		for ddIdx, dd := range drilldowns {
			if ddIdx != 0 {
				ddRow = 1 // Nested drill-downs are only supported with single stats rows.
			}
			var nextView *StatView
			var err error
			nextView, views, err = drilldownFn(viewInScope, ctx, ddRow, dd)
			if err != nil {
				return nil, fmt.Errorf("could not perform drilldown for table %q, row %d: %w", viewInScope.caption, ddRow, err)
			}
			if nextView == nil {
				// If no further drilldown table is available, stop processing the current row.
				break
			}
			viewInScope = nextView
			subTbl, err := viewInScope.fetchTableFromCSV(ctx)
			if err != nil {
				return nil, fmt.Errorf("could not fetch drilldown stats view %q for %q: %w", viewInScope.caption, dd, err)
			}
			if len(subTbl) != 1 {
				return nil, fmt.Errorf("drilldown stats view %q has %d rows, expected exactly 1", dd, len(subTbl))
			}
			for col, val := range subTbl[0] {
				// Update original table with data it does not already have.
				if _, present := row[col]; !present {
					row[col] = val
				}
			}
		}
	}
	return tbl, nil
}

// drilldown runs a 'drilldown' operation for the specified row and drilldown
// caption for the given view. It returns a stat view representing the table
// generated by the drilldown operation and an updated mapping of views since
// these can be affected by the drilldown.
func (v *StatView) drilldown(ctx context.Context, row int, dd string) (*StatView, map[string]*StatView, error) {
	var ddOpts []string
	if err := v.sess.Post(ctx, path.Join(viewOpsPath, "getAvailableDrillDownOptions"), OpArgs{v.path(), row}, &ddOpts); err != nil {
		return nil, nil, fmt.Errorf("could not retrieve drilldown options : %w", err)
	}
	var optInList bool
	for _, ddo := range ddOpts {
		if ddo == dd {
			optInList = true
			break
		}
	}
	// If the drilldown is not available, stop processing the current top-level stat row
	if !optInList {
		return nil, nil, nil
	}
	var ddCaption string
	if err := v.sess.Post(ctx, path.Join(viewOpsPath, "dodrilldownbyoption"), OpArgs{v.path(), row, dd}, &ddCaption); err != nil {
		return nil, nil, fmt.Errorf("could not perform stats drilldown: %w", err)
	}
	views, err := v.sess.Stats().Views(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not fetch updated stat views for drilldown: %w", err)
	}
	return views[ddCaption], views, nil
}

func (v *StatView) fetchTableFromCSV(ctx context.Context) (StatTable, error) {
	enableCSV := struct {
		EnableCSVLogging bool `json:"enableCsvLogging"`
	}{
		EnableCSVLogging: true,
	}
	if err := v.sess.Patch(ctx, v.path(), enableCSV); err != nil {
		return nil, fmt.Errorf("error enabling CSV logging: %w", err)
	}
	if err := v.waitForReady(ctx); err != nil {
		return nil, err
	}

	csvCount := atomic.AddUint32(&csvCounter, 1)
	csvName := strings.ReplaceAll(v.caption, " ", "_")
	csvName = fmt.Sprintf("%s_%d_%d", csvName, v.sess.ID(), csvCount)
	snapshotArgs := OpArgs{
		[]string{v.caption},
		[]string{
			fmt.Sprintf("Snapshot.View.Csv.Name: %s", csvName),
			"Snapshot.View.Contents: allPages",
			"Snapshot.View.Csv.Location: /root/.local/share/Ixia/sdmStreamManager/common",
			"Snapshot.View.Csv.GeneratingMode: kOverwriteCSVFile",
			"Snapshot.View.Csv.StringQuotes: False",
			"Snapshot.View.Csv.SupportsCSVSorting: False",
			"Snapshot.View.Csv.FormatTimestamp: True",
			"Snapshot.View.Csv.DumpTxPortLabelMap: False",
			"Snapshot.View.Csv.DecimalPrecision: 3",
		},
	}
	if err := v.sess.Post(ctx, "operations/takeviewcsvsnapshot", snapshotArgs, nil); err != nil {
		return nil, fmt.Errorf("error taking CSV snapshot: %w", err)
	}
	csvFilename := csvName + ".csv"
	defer closer.CloseAndLog(func() error {
		return v.sess.Files().Delete(ctx, csvFilename)
	}, "error deleting CSV file")
	csvBytes, err := v.sess.Files().Download(ctx, csvFilename)
	if err != nil {
		return nil, fmt.Errorf("error downloading CSV file: %w", err)
	}
	recs, err := csv.NewReader(bytes.NewReader(csvBytes)).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error parsing CSV file: %w", err)
	}

	var table StatTable
	columns := recs[0]
	for _, rec := range recs[1:] {
		row := StatRow{}
		for i := 0; i < len(rec); i++ {
			row[columns[i]] = rec[i]
		}
		table = append(table, row)
	}
	return table, nil
}

func (v *StatView) fetchEgressTableFromPages(ctx context.Context) (StatTable, error) {
	if err := v.waitForReady(ctx); err != nil {
		return nil, err
	}
	dataPath := path.Join(v.path(), "data")
	pageNum := struct {
		PageNum int `json:"currentPage"`
	}{}
	pageTable := struct {
		Columns  []string     `json:"columnCaptions"`
		Values   [][][]string `json:"pageValues"`
		NumPages int          `json:"totalPages"`
	}{}

	var table StatTable
	for i := 1; ; i++ {
		pageNum.PageNum = i
		if err := v.sess.Patch(ctx, dataPath, pageNum); err != nil {
			return nil, fmt.Errorf("error setting page num: %w", err)
		}
		if err := v.sess.Get(ctx, dataPath, &pageTable); err != nil {
			return nil, fmt.Errorf("error getting page data: %w", err)
		}
		if len(pageTable.Values) == 0 || len(pageTable.Values[0]) == 0 || len(pageTable.Values[0][0]) == 0 {
			return nil, fmt.Errorf("IxNetwork shows no stats for view %q on page %d", v.caption, i)
		}
		for _, pv := range pageTable.Values {
			for _, rec := range pv {
				row := StatRow{}
				for i := 0; i < len(rec); i++ {
					row[pageTable.Columns[i]] = rec[i]
				}
				table = append(table, row)
			}
		}
		if i == pageTable.NumPages {
			break
		}
	}
	return table, nil
}

func (v *StatView) waitForReady(ctx context.Context) error {
	const delay = 5 * time.Second
	dataPath := path.Join(v.path(), "data")
	isReady := struct {
		IsReady bool `json:"isReady"`
	}{}
	for i := 0; i < viewReadyAttempts; i++ {
		if err := v.sess.Get(ctx, dataPath, &isReady); err != nil {
			return err
		}
		if isReady.IsReady {
			return nil
		}
		sleepFn(delay)
	}
	return fmt.Errorf("timeout waiting for statistics view %q to be ready", v.caption)
}

func (v *StatView) configEgressFilters(ctx context.Context, trafficItems []string) error {
	if len(trafficItems) == 0 {
		return errors.New("No traffic items specified for egress stat view")
	}
	portFilters, err := v.fetchFilters(ctx, "availablePortFilter", nil)
	if err != nil {
		return err
	}

	// Fetch only the available traffic item filters whose name matches one of the
	// traffic items passed to this method.
	trafficItemFilters, err := v.fetchFilters(ctx, "availableTrafficItemFilter", func(name string) bool {
		for _, ti := range trafficItems {
			if ti == name {
				return true
			}
		}
		return false
	})
	if err != nil {
		return err
	}

	flowFilter := struct {
		EgressLatencyBinDisplayOption string   `json:"egressLatencyBinDisplayOption"`
		PortFilterIds                 []string `json:"portFilterIds,omitempty"`
		TrafficItemFilterIds          []string `json:"trafficItemFilterIds,omitempty"`
	}{
		EgressLatencyBinDisplayOption: "showEgressRows",
		PortFilterIds:                 portFilters,
		TrafficItemFilterIds:          trafficItemFilters,
	}
	if err := v.sess.Patch(ctx, path.Join(v.path(), "layer23TrafficFlowFilter"), flowFilter); err != nil {
		return fmt.Errorf("error configuring layer23TrafficFlowFilter: %w", err)
	}

	// Fetch all the available tracking filters but ensure they are unique.
	visited := make(map[string]bool)
	trackingFilters, err := v.fetchFilters(ctx, "availableTrackingFilter", func(name string) bool {
		if visited[name] {
			return false
		}
		visited[name] = true
		return true
	})
	if err != nil {
		return err
	}
	enumFilterPath := path.Join(v.path(), "layer23TrafficFlowFilter/enumerationFilter")
	enumFilter := struct {
		SortDirection    string `json:"sortDirection"`
		TrackingFilterID string `json:"trackingFilterId"`
	}{
		SortDirection: "ascending",
	}
	for _, f := range trackingFilters {
		enumFilter.TrackingFilterID = f
		if err := v.sess.Post(ctx, enumFilterPath, enumFilter, nil); err != nil {
			return fmt.Errorf("error adding tracking filter %q: %w", f, err)
		}
	}

	statFilters, err := v.fetchFilters(ctx, "statistic", nil)
	if err != nil {
		return err
	}
	enable := struct {
		Enabled bool `json:"enabled"`
	}{
		Enabled: true,
	}
	for _, f := range statFilters {
		if err := v.sess.Patch(ctx, f, enable); err != nil {
			return fmt.Errorf("error enabling statistic filter %q: %w", f, err)
		}
	}
	if err := v.sess.Patch(ctx, v.path(), enable); err != nil {
		return fmt.Errorf("error enabling Egress Stat View: %w", err)
	}
	return nil
}

// fetchFilters fetches the available filters at the specified relative path
// and returns a list of their href IDs. If a namePred is specified, a filter's
// name must pass the predicate for its ID to be included in the return value.
func (v *StatView) fetchFilters(ctx context.Context, relPath string, namePred func(string) bool) ([]string, error) {
	var filters []struct {
		Name  string `json:"name"`
		Links []struct {
			Href string `json:"href"`
		} `json:"links"`
	}
	if err := v.sess.Get(ctx, path.Join(v.path(), relPath), &filters); err != nil {
		return nil, fmt.Errorf("error fetching filters at this %q: %w", relPath, err)
	}
	var hrefs []string
	for _, f := range filters {
		if len(f.Links) > 0 && (namePred == nil || namePred(f.Name)) {
			hrefs = append(hrefs, f.Links[0].Href)
		}
	}
	return hrefs, nil
}
