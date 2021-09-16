package Stratum

import (
	"errors"
)

type SetDifficultyParams struct {
	Difficulty Difficulty
}

func (p *SetDifficultyParams) Read(n *notification) error {
	if len(n.params) != 1 {
		return errors.New("Incorrect parameter length")
	}

	if !ValidDifficulty(n.params[0]) {
		return errors.New("Invalid difficulty")
	}

	p.Difficulty = n.params[0]

	return nil
}

func SetDifficulty(d Difficulty) notification {
	return Notification(MiningSetDifficulty, []interface{}{d})
}
