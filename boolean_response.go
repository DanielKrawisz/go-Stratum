package Stratum

import (
	"errors"
)

type BooleanResult struct {
	Result bool
}

func (b *BooleanResult) Read(r *response) error {
	var ok bool
	b.Result, ok = r.result.(bool)
	if !ok {
		return errors.New("Invalid value")
	}

	return nil
}

func BooleanResponse(id MessageID, x bool) response {
	return Response(id, x)
}
