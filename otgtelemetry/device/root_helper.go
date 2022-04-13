package device

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/otgtelemetry"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// WithReplica adds the replica number to the context metadata of the gNMI
// server query.
func (n *DevicePath) WithReplica(replica int) *DevicePath {
	genutil.PutReplica(n, replica)
	return n
}

// WithSubscriptionMode specifies the subscription mode in the underlying gNMI
// subscribe.
func (n *DevicePath) WithSubscriptionMode(mode gpb.SubscriptionMode) *DevicePath {
	genutil.PutSubscriptionMode(n, mode)
	return n
}

// WithClient allows the user to provide a gNMI client. This allows for creation
// of tests for multiple gNMI clients to a single DUT.
func (n *DevicePath) WithClient(c gpb.GNMIClient) *DevicePath {
	genutil.PutClient(n, c)
	return n
}

// NewBatch returns a newly instantiated SetRequestBatch object for batching set requests.
func (d *DevicePath) NewBatch() *oc.Batch {
	return oc.NewBatch(d)
}
