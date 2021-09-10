package telemgo

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

const (
	metadataKeyPrefix   = "metadata-"
	subscriptionModeKey = "subscriptionMode"
	clientKey           = "client"
)

// PutClient sets the client as metadata request option.
func PutClient(n rootPath, client gpb.GNMIClient) {
	n.PutCustomData(clientKey, client)
}

// PutReplica sets the replica number as a metadata request option.
func PutReplica(n rootPath, replica int) {
	n.PutCustomData(metadataKeyPrefix+"replica", strconv.Itoa(replica))
}

// PutSubscriptionMode sets the subscription mode as a request option.
func PutSubscriptionMode(n rootPath, subMode gpb.SubscriptionMode) {
	n.PutCustomData(subscriptionModeKey, subMode)
}

type rootPath interface {
	PutCustomData(key string, val interface{})
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
			return nil, errors.Errorf("CustomData key %q but value is not SubscriptionMode type (%T, %v)", subscriptionModeKey, v, v)
		}
		opts.subMode = m
	}
	if v, ok := customData[clientKey]; ok {
		if v == nil {
			return nil, errors.Errorf("CustomData key %q but value is nil", clientKey)
		}
		m, ok := v.(gpb.GNMIClient)
		if !ok {
			return nil, errors.Errorf("CustomData key %q but value is not GNMIClient type (%T, %v)", clientKey, v, v)
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
			return nil, errors.Errorf("CustomData metadata key %q but value is not string type (%T, %v)", k, v, v)
		}
		md[strings.TrimPrefix(k, metadataKeyPrefix)] = sv
	}
	opts.md = metadata.New(md)
	return opts, nil
}
