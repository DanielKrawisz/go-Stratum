package Stratum

import (
	"encoding/hex"
	"errors"
)

type NotifyParams struct {
	JobID         ID
	Digest        []byte
	GenerationTX1 []byte
	GenerationTX2 []byte
	Path          [][]byte
	Version       uint32
	Target        []byte
	Timestamp     uint32
	Clean         bool
}

func (p *NotifyParams) read(n *notification) error {
	if len(n.params) != 9 {
		return errors.New("invalid format")
	}

	jobID, ok := n.params[0].(string)
	if !ok {
		return errors.New("invalid format")
	}

	digest, ok := n.params[1].(string)
	if !ok {
		return errors.New("invalid format")
	}

	gtx1, ok := n.params[2].(string)
	if !ok {
		return errors.New("invalid format")
	}

	gtx2, ok := n.params[3].(string)
	if !ok {
		return errors.New("invalid format")
	}

	path, ok := n.params[4].([]string)
	if !ok {
		return errors.New("invalid format")
	}

	version, ok := n.params[5].(string)
	if !ok {
		return errors.New("invalid format")
	}

	target, ok := n.params[6].(string)
	if !ok {
		return errors.New("invalid format")
	}

	time, ok := n.params[7].(string)
	if !ok {
		return errors.New("invalid format")
	}

	p.Clean, ok = n.params[8].(bool)
	if !ok {
		return errors.New("invalid format")
	}

	var err error
	p.Digest, err = hex.DecodeString(digest)
	if err != nil || len(p.Digest) != 32 {
		return errors.New("invalid format")
	}

	p.Target, err = hex.DecodeString(target)
	if err != nil || len(p.Digest) != 4 {
		return errors.New("invalid format")
	}

	p.JobID, err = decodeID(jobID)
	if err != nil {
		return errors.New("invalid format")
	}

	p.GenerationTX1, err = hex.DecodeString(gtx1)
	if err != nil {
		return errors.New("invalid format")
	}

	p.GenerationTX2, err = hex.DecodeString(gtx2)
	if err != nil {
		return errors.New("invalid format")
	}

	p.Version, err = decodeBigEndian(version)
	if err != nil {
		return errors.New("invalid format")
	}

	p.Timestamp, err = decodeBigEndian(time)
	if err != nil {
		return errors.New("invalid format")
	}

	p.Path = make([][]byte, len(path))
	for i := 0; i < len(path); i++ {
		p.Path[i], err = hex.DecodeString(path[i])
		if err != nil || len(p.Digest) != 32 {
			return errors.New("invalid format")
		}
	}

	return nil
}

func Notify(n NotifyParams) notification {
	params := make([]interface{}, 9)

	params[0] = encodeID(n.JobID)
	params[1] = hex.EncodeToString(n.Digest)
	params[2] = hex.EncodeToString(n.GenerationTX1)
	params[3] = hex.EncodeToString(n.GenerationTX2)

	path := make([]string, len(n.Path))
	for i := 0; i < len(n.Path); i++ {
		path[i] = hex.EncodeToString(n.Path[i])
	}

	params[4] = path
	params[5] = encodeBigEndian(n.Version)
	params[6] = hex.EncodeToString(n.Target)
	params[7] = encodeBigEndian(n.Timestamp)
	params[8] = n.Clean

	return Notification(MiningNotify, params)
}
