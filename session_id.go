package Stratum

// A stratum session id is assigned by the mining pool to a miner and it
// is included in the coinbase script of the block that is produced.
type SessionID uint32

func EncodeSessionID(id SessionID) string {}

func DecodeSessionID(string) (SessionID, error) {}
