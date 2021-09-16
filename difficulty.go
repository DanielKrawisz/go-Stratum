package Stratum

type Difficulty interface{}

// Difficulty can be given as a uint or a float.
func ValidDifficulty(u Difficulty) bool {
	switch d := u.(type) {
	case uint64:
		return d > 0
	case float64:
		return d > 0
	default:
		return false
	}
}
