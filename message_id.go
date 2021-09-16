package Stratum

// MessageID is a unique identifier that is different for each notification
// and request / response.
type MessageID interface{}

// MessageIDs are allowed to be integers or strings.
func ValidMessageID(id MessageID) bool {
	switch id.(type) {
	case uint64:
		return true
	case string:
		return true
	default:
		return false
	}
}
