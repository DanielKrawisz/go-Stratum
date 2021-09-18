package Stratum

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
)

// A stratum session id is assigned by the mining pool to a miner and it
// is included in the coinbase script of the block that is produced.
type SessionID uint32

func EncodeSessionID(id SessionID) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return hex.EncodeToString(b)
}

func DecodeSessionID(s string) (SessionID, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}

	if len(b) != 4 {
		return 0, errors.New("Invalid format")
	}

	var x uint32
	binary.Read(bytes.NewBuffer(b), binary.BigEndian, &x)
	return SessionID(x), nil
}
