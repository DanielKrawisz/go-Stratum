package Stratum

import "encoding/json"

// MessageID is a unique identifier that is different for each notification
// and request / response.
type MessageID []byte

// MessageIDs are allowed to be integers or strings.
func MessageIDFromUInt(x uint64) MessageID {
	id, _ := json.Marshal(x)
	return id
}

func MessageIDFromString(x string) MessageID {
	id, _ := json.Marshal(x)
	return id
}
