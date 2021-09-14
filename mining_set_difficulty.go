package Stratum

type setDifficulty struct {
	Notification
}

func (au *setDifficulty) Params() uint64 {}

func NewSetDifficulty(RequestID, uint64) (*setDifficulty, error) {}
