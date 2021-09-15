package Stratum

type submitRequest struct {
	Request
}

func (r *submitRequest) Params() (Share, error) {}

func NewSubscribeRequest(MessageID, Share) *submitRequest {}

type submitResponse booleanResponse
