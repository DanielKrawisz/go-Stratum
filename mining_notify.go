package Stratum

type NotifyParams struct {
	JobID         ID
	Digest        [32]byte
	GenerationTX1 []byte
	GenerationTX2 []byte
	Path          [][32]byte
	Version       uint32
	Target        []byte
	Timestamp     uint32
}

func (p *NotifyParams) read(n *notification) error {}

func Notify(NotifyParams) notification {}
