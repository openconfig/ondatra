package telemgo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/metadata"
	"github.com/openconfig/gnmi/errdiff"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

func TestExtractRequestOptions(t *testing.T) {
	c := gpb.NewGNMIClient(nil)
	tests := []struct {
		name          string
		inCustomData  map[string]interface{}
		want          *requestOpts
		wantErrSubstr string
	}{{
		name: "invalid client mode",
		inCustomData: map[string]interface{}{
			clientKey: 1,
		},
		wantErrSubstr: "value is not GNMIClient type",
	}, {
		name: "nil client mode",
		inCustomData: map[string]interface{}{
			clientKey: nil,
		},
		wantErrSubstr: "value is nil",
	}, {
		name: "get client mode",
		inCustomData: map[string]interface{}{
			clientKey: c,
		},
		want: &requestOpts{client: c, md: metadata.MD{}},
	}, {
		name: "get subscription mode",
		inCustomData: map[string]interface{}{
			subscriptionModeKey: gpb.SubscriptionMode_ON_CHANGE,
		},
		want: &requestOpts{subMode: gpb.SubscriptionMode_ON_CHANGE, md: metadata.MD{}},
	}, {
		name: "single metadata field",
		inCustomData: map[string]interface{}{
			metadataKeyPrefix + "replica": "5",
		},
		want: &requestOpts{md: metadata.Pairs("replica", "5")},
	}, {
		name: "single metadata field but not string type",
		inCustomData: map[string]interface{}{
			metadataKeyPrefix + "replica": 5,
		},
		wantErrSubstr: "value is not string type",
	}, {
		name: "non-metadata field, non-string type",
		inCustomData: map[string]interface{}{
			"replica": 5,
		},
		want: &requestOpts{md: metadata.MD{}},
	}, {
		name: "multiple metadata fields",
		inCustomData: map[string]interface{}{
			metadataKeyPrefix + "replica": "5",
			metadataKeyPrefix + "foo":     "bar",
		},
		want: &requestOpts{md: metadata.Pairs("replica", "5", "foo", "bar")},
	}, {
		name: "mix of metadata and non-metadata fields",
		inCustomData: map[string]interface{}{
			metadataKeyPrefix + "replica": "5",
			metadataKeyPrefix + "foo":     "bar",
			"replica":                     10,
			"foo":                         uint64(1234),
		},
		want: &requestOpts{md: metadata.Pairs("replica", "5", "foo", "bar")},
	}, {
		name: "mix of metadata and non-metadata fields with a non-string metadata type",
		inCustomData: map[string]interface{}{
			metadataKeyPrefix + "replica": "5",
			metadataKeyPrefix + "foo":     uint64(2345),
			"replica":                     10,
			"foo":                         uint64(1234),
		},
		wantErrSubstr: "value is not string type",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractRequestOpts(tt.inCustomData)
			switch err {
			case nil:
				if diff := cmp.Diff(tt.want, got, cmp.Comparer(func(a, b gpb.GNMIClient) bool {
					return a == b
				}),
					cmp.AllowUnexported(requestOpts{})); diff != "" {
					t.Errorf("extractMetadata: input struct after unmarshalling does not match (-want +got):\n%s", diff)
				}
			default:
				if diff := errdiff.Substring(err, tt.wantErrSubstr); diff != "" {
					t.Fatalf("extractMetadata: did not get expected error substring:\n%s", diff)
				}
			}
		})
	}
}
