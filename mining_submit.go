package Stratum

type SubmitParams Share

func Submit(MessageID, SubmitParams) request {}

type SubmitResult BooleanResult

func SubmitResponse(id MessageID, b bool) response {
	return BooleanResponse(id, b)
}
