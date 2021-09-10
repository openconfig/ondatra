// Package stats provides common helper functions that are associated
// with ONDATRA return types to calculate statistics for telemetry values.
package stats

import (
	"sort"
	"testing"

	"github.com/openconfig/ondatra/telemetry"
)

// MeanRateUint64 calculates the average rate across the supplied interval
// from the values supplied in the vals slice. It sorts the values by
// timestamp and uses the earliest and latest to calculate the rate. Errors
// are reported using the supplied testing.TB.
func MeanRateUint64(t testing.TB, vals []*telemetry.QualifiedUint64) float64 {
	t.Helper()

	if len(vals) < 2 {
		t.Fatalf("cannot calculate rate from %d length slice, got: %v", len(vals), vals)
	}

	sort.Slice(vals, func(i, j int) bool {
		return vals[i].Timestamp.Before(vals[j].Timestamp)
	})

	return float64(vals[len(vals)-1].Val(t)-vals[0].Val(t)) / vals[len(vals)-1].Timestamp.Sub(vals[0].Timestamp).Seconds()
}
