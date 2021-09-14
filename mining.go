package Stratum

type JobId uint32

type WorkerName string

type Worker struct{}

func (w *Worker) MarshallJSON() ([]byte, error) {}

func (w *Worker) UnmarshallJSON([]byte) error {}

// A share is the data returned by the worker. Job + Share = Proof
type Share struct {
	Time               uint32
	Nonce              uint32
	ExtraNonce2        uint64
	GeneralPurposeBits *uint32
}

func MakeShare(time uint32, nonce uint32, extraNonce2 uint64) Share {
	return Share{
		Time:               time,
		Nonce:              nonce,
		ExtraNonce2:        extraNonce2,
		GeneralPurposeBits: nil}
}

func MakeShareASICBoost(time uint32, nonce uint32, extraNonce2 uint64, gpb uint32) Share {
	bits := new(uint32)
	*bits = gpb
	return Share{
		Time:               time,
		Nonce:              nonce,
		ExtraNonce2:        extraNonce2,
		GeneralPurposeBits: bits}
}

func (x *Share) MarshallJSON() ([]byte, error) {}

func (x *Share) UnmarshallJSON([]byte) error {}
