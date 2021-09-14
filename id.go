package Stratum

// A stratum session id is assigned by the mining pool to a miner and it
// is included in the coinbase script of the block that is produced.
type ID uint32

func EncodeId(id ID) (string, error) {}

func DecodeId(string) (ID, error) {}
