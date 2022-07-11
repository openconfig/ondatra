// Copyright 2022 Google LLC
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

package httpovergrpc

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	// Measures
	mNumRequests  = stats.Int64("num_requests", "number of requests received.", "1")
	mProxyErrors  = stats.Int64("num_proxy_errors", "number of proxy errors", "1")
	mClientErrors = stats.Int64("num_client_errors", "number of errors that originate from end client", "1")

	// Tag Keys
	urlKey, _   = tag.NewKey("url")
	errorKey, _ = tag.NewKey("errorDesc")

	// Views
	numRequestsView = &view.View{
		Name:        "num_requests",
		Description: "total number of requests received",
		Measure:     mNumRequests,
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{urlKey},
	}
	proxyErrorsView = &view.View{
		Name:        "num_proxy_errors",
		Description: "number of proxy errors",
		Measure:     mProxyErrors,
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{urlKey, errorKey},
	}
	clientErrorsView = &view.View{
		Name:        "num_client_errors",
		Description: "number of errors that originate from end client",
		Measure:     mClientErrors,
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{urlKey, errorKey},
	}
)

// RegisterMetricViews registers all metric views for proxy for Opencensus to actively record.
func RegisterMetricViews() error {
	return view.Register(numRequestsView, proxyErrorsView, clientErrorsView)
}
