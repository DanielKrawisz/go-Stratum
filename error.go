package Stratum

type ErrorCode uint32

// incomplete
const (
	None ErrorCode = iota
)

type Error struct {
	Code    ErrorCode
	Message string
}

func (e *Error) MarshallJSON() ([]byte, error) {}

func (e *Error) UnmarshallJSON([]byte) error {}
