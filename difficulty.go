package Stratum

type Difficulty []byte

// Difficulty can be given as a uint or a float.
func DifficultyFromUInt(uint64) {}

func DifficultyFromFloat(float64) {}

func DecodeDifficulty(Difficulty) (float64, error) {}
