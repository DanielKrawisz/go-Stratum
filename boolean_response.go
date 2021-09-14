package Stratum

type BooleanResponse struct {
	Response
}

func (b *BooleanResponse) Valid() bool {}

func (b *BooleanResponse) Result() bool {}
