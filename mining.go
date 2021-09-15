package Stratum

type JobID uint32

type WorkerName string

// Worker represents a miner who is doing work for the pool.
// This would be used in an implementation of a Stratum server and is
// not part of the Stratum protocol.
type Worker struct {
	Name WorkerName
	ID SessionID
	uint32 ExtraNonce2Size
	*uint32 VersionMask
}

// A share is the data returned by the worker in mining.submit.
type Share struct {
	Name  WorkerName
	JobID JobID
	work.Share
}

func (s *Share) MarshallJSON() ([]byte, error) {}

func (s *Share) UnmarshallJSON([]byte) error {}
