package Stratum

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
)

// A stratum session id is assigned by the mining pool to a miner and it
// is included in the coinbase script of the block that is produced.
// we also use ID for job ids.
type ID uint32

func encodeBigEndian(id uint32) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return hex.EncodeToString(b)
}

func decodeBigEndian(s string) (uint32, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}

	if len(b) != 4 {
		return 0, errors.New("Invalid format")
	}

	var x uint32
	binary.Read(bytes.NewBuffer(b), binary.BigEndian, &x)
	return x, nil
}

func encodeLittleEndian(id uint32) string {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(id))
	return hex.EncodeToString(b)
}

func decodeLittleEndian(s string) (uint32, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}

	if len(b) != 4 {
		return 0, errors.New("Invalid format")
	}

	var x uint32
	binary.Read(bytes.NewBuffer(b), binary.LittleEndian, &x)
	return x, nil
}

func encodeID(id ID) string {
	return encodeBigEndian(uint32(id))
}

func decodeID(s string) (ID, error) {
	x, err := decodeBigEndian(s)
	if err != nil {
		return 0, err
	}

	return ID(x), nil
}
