package Stratum

import (
	"errors"
)

type Extension uint8

const (
	Unknown = iota
	VersionRolling
	MinimumDifficulty
	SubscribeExtranonce
	Info
)

func EncodeExtension(m Extension) (string, error) {
	switch m {
	case VersionRolling:
		return "version_rolling", nil
	case MinimumDifficulty:
		return "minimum_difficulty", nil
	case SubscribeExtranonce:
		return "subscribe_extranonce", nil
	case Info:
		return "info", nil
	default:
		return "", errors.New("Unkown Stratum extension")
	}
}

func DecodeExtension(m string) (Extension, error) {
	switch m {
	case "version_rolling":
		return VersionRolling, nil
	case "minimum_difficulty":
		return MinimumDifficulty, nil
	case "subscribe_extranonce":
		return SubscribeExtranonce, nil
	case "info":
		return Info, nil
	default:
		return Unknown, errors.New("Unkown Stratum extension")
	}
}
