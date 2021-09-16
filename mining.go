package Stratum

import "github.com/DanielKrawisz/go-work"

type JobID uint32

type WorkerName string

// Worker represents a miner who is doing work for the pool.
// This would be used in an implementation of a Stratum server and is
// not part of the Stratum protocol.
type Worker struct {
	Name            WorkerName
	ID              SessionID
	ExtraNonce2Size uint32
	VersionMask     *uint32
}

// A share is the data returned by the worker in mining.submit.
type Share struct {
	Name  WorkerName
	JobID JobID
	work.Share
}

func (p *Share) read(r *request) error {}
