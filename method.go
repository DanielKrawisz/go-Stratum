package Stratum

import (
	"errors"
)

type Method uint8

const (
	Unset Method = iota
	MiningAuthorize
	MiningConfigure
	MiningSubscribe
	MiningNotify
	MiningSubmit
	MiningSetDifficulty
	MiningSetVersionMask
	MiningSetExtraNonce
	MiningSuggestDifficulty
	MiningSuggestTarget
	ClientGetVersion
	ClientReconnect
	ClientGetTransactions
	ClientShowMessage
)

func EncodeMethod(m Method) (string, error) {
	switch m {
	case MiningAuthorize:
		return "mining.authorize", nil
	case MiningConfigure:
		return "mining.configure", nil
	case MiningSubscribe:
		return "mining.subscribe", nil
	case MiningNotify:
		return "mining.notify", nil
	case MiningSubmit:
		return "mining.submit", nil
	case MiningSetDifficulty:
		return "mining.set_difficulty", nil
	case MiningSetVersionMask:
		return "mining.set_version_mask", nil
	case MiningSetExtraNonce:
		return "mining.set_extra_nonce", nil
	case MiningSuggestDifficulty:
		return "mining.suggest_difficulty", nil
	case MiningSuggestTarget:
		return "mining.suggest_target", nil
	case ClientGetVersion:
		return "mining.get_version", nil
	case ClientReconnect:
		return "mining.reconnect", nil
	case ClientGetTransactions:
		return "mining.transactions", nil
	case ClientShowMessage:
		return "mining.show_message", nil
	default:
		return "", errors.New("Unkown Stratum method")
	}
}

func DecodeMethod(m string) (Method, error) {
	switch m {
	case "mining.authorize":
		return MiningAuthorize, nil
	case "mining.configure":
		return MiningConfigure, nil
	case "mining.subscribe":
		return MiningSubscribe, nil
	case "mining.notify":
		return MiningNotify, nil
	case "mining.submit":
		return MiningSubmit, nil
	case "mining.set_difficulty":
		return MiningSetDifficulty, nil
	case "mining.set_version_mask":
		return MiningSetVersionMask, nil
	case "mining.set_extra_nonce":
		return MiningSetExtraNonce, nil
	case "mining.suggest_difficulty":
		return MiningSuggestDifficulty, nil
	case "mining.suggest_target":
		return MiningSuggestTarget, nil
	case "mining.get_version":
		return ClientGetVersion, nil
	case "mining.reconnect":
		return ClientReconnect, nil
	case "mining.transactions":
		return ClientGetTransactions, nil
	case "mining.show_message":
		return ClientShowMessage, nil
	default:
		return Unset, errors.New("Unkown Stratum method")
	}
}
