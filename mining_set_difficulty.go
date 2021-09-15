package Stratum

type setDifficulty struct {
	Notification
}

func (au *setDifficulty) Params() (float64, error) {}

func NewSetDifficulty(MethodID, Difficulty) *setDifficulty {}
