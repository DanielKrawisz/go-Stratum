package Stratum

import "encoding/json"

type Difficulty []byte

// Difficulty can be given as a uint or a float.
func DifficultyFromUInt(x uint64) Difficulty {
	id, _ := json.Marshal(x)
	return id
}

func DifficultyFromFloat(x float64) Difficulty {
	id, _ := json.Marshal(x)
	return id
}
