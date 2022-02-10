// Copyright 2021 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ixconfig

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Format: /multivalue[@source = '<path to parent> <object name>']
	multivalueRe = regexp.MustCompile(`^/multivalue\[@source = '(\S+) (\w+)'\]$`)
	// Format: /path/to/parent/<objectName> [@alias= '<alias>']
	aliasRe = regexp.MustCompile(`^(.*/)(\w+)\[@alias = '(.+)'\]$`)
	// Format: /path/to/parent/<objectName>[<index>]
	arrayRe = regexp.MustCompile(`^(.*/)(\w+)\[([1-9][0-9]*)\]$`)
	// Format: /path/to/parent/<objectName>
	singleRe = regexp.MustCompile(`^(.*/)(\w+)$`)

	excludeNodesFromIndexing = map[string]bool{
		"dynamicUpdate":            true,
		"tracking":                 true,
		"transmissionDistribution": true,
	}
)

// XPath represents the specific xmlpath (XPath) indexing for Ixia IxNetwork API JSON config nodes.
type XPath struct {
	// The full XPath to the parent node (as a string).
	parentXPath string
	// Name of the object represented by the XPath.
	objectName string
	// Alias of the object, used as an alternative to numerical indexing.
	// If this is set, it will override numerical indexing.
	alias string
	// Whether the object is of the Multivalue type.
	isMultivalue bool
	// XPath indices for arrays begin at '1', not '0'. If this value is 0, we
	// assume that this XPath does not refer to an object within an array.
	index uint64
}

func (xp *XPath) copyXPath() *XPath {
	if xp == nil {
		return nil
	}
	return &XPath{
		parentXPath:  xp.parentXPath,
		objectName:   xp.objectName,
		alias:        xp.alias,
		isMultivalue: xp.isMultivalue,
		index:        xp.index,
	}
}

// String returns the XPath as would appear in an IxNetwork JSON config.
func (xp *XPath) String() string {
	if xp.isMultivalue {
		return fmt.Sprintf("/multivalue[@source = '%s %s']", path.Clean(xp.parentXPath), xp.objectName)
	}
	var idxSuffix string

	// There are some array objects whose XPaths are not indexed in exported configs. As an example,
	// instead of the expected XPath '/traffic/trafficItem[1]/transmissionDistribution[1]',
	// we instead see '/traffic/trafficItem[1]/transmissionDistribution'.
	// We specifically exclude these objects when generating indexed XPaths.
	if !strings.HasPrefix(xp.parentXPath, "/globals") && !excludeNodesFromIndexing[xp.objectName] {
		if xp.alias != "" {
			idxSuffix = fmt.Sprintf("[@alias = '%s']", xp.alias)
		} else if xp.index != 0 {
			idxSuffix = fmt.Sprintf("[%d]", xp.index)
		}
	}
	return path.Join(xp.parentXPath, fmt.Sprintf("%s%s", xp.objectName, idxSuffix))
}

// MarshalJSON implements the encoding/json.Marshaler interface.
func (xp *XPath) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(xp.String())), nil
}

// UnmarshalJSON fills in an XPath object from an XPath string, implementing the encoding/json.Marshaler interface.
// The provided XPath object is assumed to have all fields set to their zero value.
func (xp *XPath) UnmarshalJSON(b []byte) error {
	if b == nil {
		return nil
	}
	// XPath strings are assumed to always start with a "/" and equivalent to their path.Clean representation.
	str, err := strconv.Unquote(string(b))
	if err != nil {
		return fmt.Errorf("Could not interpret %q as a JSON string: %v", string(b), err)
	}
	return xp.unmarshalString(str)
}

func (xp *XPath) unmarshalString(str string) error {
	if str == "/" {
		// We are at the root of the hierarchy, no objectName to set.
		xp.parentXPath = "/"
		return nil
	}
	matches := multivalueRe.FindStringSubmatch(str)
	if matches != nil {
		xp.parentXPath = path.Clean(matches[1])
		xp.objectName = matches[2]
		xp.isMultivalue = true
		return nil
	}
	matches = aliasRe.FindStringSubmatch(str)
	if matches != nil {
		xp.parentXPath = path.Clean(matches[1])
		xp.objectName = matches[2]
		xp.alias = matches[3]
		return nil
	}
	matches = arrayRe.FindStringSubmatch(str)
	if matches != nil {
		xp.parentXPath = path.Clean(matches[1])
		xp.objectName = matches[2]
		var err error
		xp.index, err = strconv.ParseUint(matches[3], 10, 64)
		if err != nil {
			// Should never happen because the regex can only match on an integer index.
			return fmt.Errorf("could not parse numerical index from XPath string %q: %v", str, err)
		}
		return nil
	}
	matches = singleRe.FindStringSubmatch(str)
	if matches != nil {
		xp.parentXPath = path.Clean(matches[1])
		xp.objectName = matches[2]
		return nil
	}
	return fmt.Errorf("could not parse %q as an IxNetwork XPath", str)
}

// ParseXPath parses the specified string to an XPath.
func ParseXPath(str string) (*XPath, error) {
	xp := &XPath{}
	if err := xp.unmarshalString(str); err != nil {
		return nil, err
	}
	return xp, nil
}
