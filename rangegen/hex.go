package rangegen

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/pkg/errors"
)

const sysIDLen = 6

var maxSysID = maxVal(sysIDLen)

// SystemIDs returns a channel of the specified number of hex-formatted system IDs.
func SystemIDs(startID string, count uint32) (<-chan string, error) {
	const strLen = len("ff ff ff ff ff ff")

	b, err := hex.DecodeString(strings.NewReplacer(":", "", ".", "", " ", "").Replace(startID))
	if err != nil {
		return nil, errors.Wrapf(err, "could not decode system ID %q as hex", startID)
	}
	if len(b) > sysIDLen {
		return nil, errors.Errorf("system ID %q is too large (max size is 6 bytes)", startID)
	}

	id := new(big.Int).SetBytes(b)
	incrCount := big.NewInt(int64(count) - 1)
	last := incrCount.Add(id, incrCount)
	if last.Cmp(maxSysID) == 1 {
		return nil, errors.New("requested ID range exceeds max system ID value")
	}

	idBytes := make([]byte, sysIDLen)
	strBytes := make([]byte, strLen)
	idCh := make(chan string, 100)
	go func() {
		defer close(idCh)
		const hexDigit = "0123456789abcdef"
		for i := 0; i < int(count); i++ {
			id.FillBytes(idBytes)
			strBytes = strBytes[:0]
			for j, idb := range idBytes {
				if j != 0 {
					strBytes = append(strBytes, ' ')
				}
				// Append characters, using first/last 4 bits separately to index into hexDigits.
				strBytes = append(strBytes, hexDigit[uint8(idb)>>4], hexDigit[uint8(idb&0xf)])
			}
			id.Add(id, one)
			idCh <- string(strBytes)
		}
	}()
	return idCh, nil
}
