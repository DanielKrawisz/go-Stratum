package Stratum

import "errors"

type setDifficulty struct {
	Notification
}

func (au *setDifficulty) Params() (float64, error) {
	if len(au.params) != 1 {
		return 0, errors.New("Incorrect parameter length")
	}

	var difficulty float64
	err := json.Unmarshall(au.params[0], difficulty)
	if err != nil {
		return 0, err
	}

	if difficulty <= 0 {
		return 0, errors.New("Invalid difficulty value; must be >= 0")
	}

	return difficulty, nil
}

func NewSetDifficulty(d Difficulty) *setDifficulty {
	return &setDifficulty{
		Method: EncodeMethod(MiningSetDifficulty),
		Params: d,
	}
}
