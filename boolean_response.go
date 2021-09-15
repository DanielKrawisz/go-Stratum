package Stratum

type booleanResponse struct {
	Response
}

func (b *booleanResponse) Result() (bool, error) {}

func NewBooleanResponse(bool) *booleanResponse {}
