package Stratum

import (
	"encoding/hex"
	"errors"

	"github.com/DanielKrawisz/go-work"
)

type WorkerName string

// Worker represents a miner who is doing work for the pool.
// This would be used in an implementation of a Stratum server and is
// not part of the Stratum protocol.
type Worker struct {
	Name            WorkerName
	SessionID       ID
	ExtraNonce2Size uint32
	VersionMask     *uint32
}

// A share is the data returned by the worker in mining.submit.
type Share struct {
	Name  WorkerName
	JobID ID
	work.Share
}

func (p *Share) read(r *request) error {
	if len(r.params) < 5 || len(r.params) > 6 {
		return errors.New("Invalid format")
	}

	name, ok := r.params[0].(string)
	if !ok {
		return errors.New("Invalid format")
	}

	jobID, ok := r.params[1].(string)
	if !ok {
		return errors.New("Invalid format")
	}

	extraNonce2, ok := r.params[2].(string)
	if !ok {
		return errors.New("Invalid format")
	}

	time, ok := r.params[3].(string)
	if !ok {
		return errors.New("Invalid format")
	}

	nonce, ok := r.params[4].(string)
	if !ok {
		return errors.New("Invalid format")
	}

	if len(r.params) == 6 {
		gpr, ok := r.params[5].(string)
		if !ok {
			return errors.New("Invalid format")
		}

		GPR, err := decodeLittleEndian(gpr)
		if err != nil {
			return err
		}

		p.GeneralPurposeBits = &GPR
	}

	var err error
	p.Nonce, err = decodeBigEndian(nonce)
	if err != nil {
		return err
	}

	p.Time, err = decodeBigEndian(time)
	if err != nil {
		return err
	}

	p.ExtraNonce2, err = hex.DecodeString(extraNonce2)
	if err != nil {
		return err
	}

	p.JobID, err = decodeID(jobID)
	if err != nil {
		return err
	}

	p.Name = WorkerName(name)
	return nil
}
