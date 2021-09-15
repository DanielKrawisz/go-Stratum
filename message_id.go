package Stratum

// MessageID is a unique identifier that is different for each notification
// and request / response.
type MessageID []byte

// MessageIDs are allowed to be integers or strings.
func MessageIDFromUInt(uint64) {}

func MessageIDFromString(string) {}
