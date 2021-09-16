package Stratum

type ErrorCode uint32

// the Stratum protocol does not define any error codes. Each pool
// has its own set of errors, apparently. You can define your own.
const (
	None ErrorCode = iota
)

// Error is a 2-element json array.
type Error struct {
	Code    ErrorCode
	Message string
}
