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

package genutil

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc/metadata"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

const (
	metadataKeyPrefix   = "metadata-"
	subscriptionModeKey = "subscriptionMode"
	clientKey           = "client"
)

// PutClient sets the client as metadata request option.
func PutClient(n FakeRootPathStruct, client gpb.GNMIClient) {
	n.PutCustomData(clientKey, client)
}

// PutReplica sets the replica number as a metadata request option.
func PutReplica(n FakeRootPathStruct, replica int) {
	n.PutCustomData(metadataKeyPrefix+"replica", strconv.Itoa(replica))
}

// PutSubscriptionMode sets the subscription mode as a request option.
func PutSubscriptionMode(n FakeRootPathStruct, subMode gpb.SubscriptionMode) {
	n.PutCustomData(subscriptionModeKey, subMode)
}

type requestOpts struct {
	subMode gpb.SubscriptionMode
	client  gpb.GNMIClient
	md      metadata.MD
}

// extractRequestOpts translates the root path's custom data to request options.
func extractRequestOpts(customData map[string]interface{}) (*requestOpts, error) {
	opts := new(requestOpts)
	if v, ok := customData[subscriptionModeKey]; ok {
		m, ok := v.(gpb.SubscriptionMode)
		if !ok {
			return nil, fmt.Errorf("customData key %q but value is not SubscriptionMode type (%T, %v)", subscriptionModeKey, v, v)
		}
		opts.subMode = m
	}
	if v, ok := customData[clientKey]; ok {
		if v == nil {
			return nil, fmt.Errorf("customData key %q but value is nil", clientKey)
		}
		m, ok := v.(gpb.GNMIClient)
		if !ok {
			return nil, fmt.Errorf("customData key %q but value is not GNMIClient type (%T, %v)", clientKey, v, v)
		}
		opts.client = m
	}
	md := make(map[string]string)
	for k, v := range customData {
		if !strings.HasPrefix(k, metadataKeyPrefix) {
			continue
		}
		sv, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("customData metadata key %q but value is not string type (%T, %v)", k, v, v)
		}
		md[strings.TrimPrefix(k, metadataKeyPrefix)] = sv
	}
	opts.md = metadata.New(md)
	return opts, nil
}
