package rangegen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSystemIDs(t *testing.T) {
	tests := []struct {
		desc    string
		startID string
		count   uint32
		wantIDs []string
		wantErr bool
	}{{
		desc:    "Invalid system ID",
		startID: "invalid",
		count:   0,
		wantErr: true,
	}, {
		desc:    "System ID too large",
		startID: "11 22 33 44 55 66 77",
		count:   0,
		wantErr: true,
	}, {
		desc:    "System ID overflow",
		startID: "ff ff ff ff ff ff",
		count:   2,
		wantErr: true,
	}, {
		desc:    "Successful iteration",
		startID: "aa aa aa aa aa ff",
		count:   2,
		wantIDs: []string{"aa aa aa aa aa ff", "aa aa aa aa ab 00"},
	}, {
		desc:    "Successful iteration, different but valid format",
		startID: "1234.5678.9abc",
		count:   2,
		wantIDs: []string{"12 34 56 78 9a bc", "12 34 56 78 9a bd"},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			idsCh, gotErr := SystemIDs(test.startID, test.count)
			if (gotErr != nil) != test.wantErr {
				t.Fatalf("SystemIDs{%s, %d}: unexpected error result, got err: %v, want err? %t",
					test.startID, test.count, gotErr, test.wantErr)
			}
			if test.wantErr {
				return
			}

			gotIDs := make([]string, 0)
			for id := range idsCh {
				gotIDs = append(gotIDs, id)
			}
			if diff := cmp.Diff(test.wantIDs, gotIDs); diff != "" {
				t.Fatalf("SystemIDs{%s, %d}: unexpected diff in generated list of IDs: %s",
					test.startID, test.count, diff)
			}
		})
	}
}
