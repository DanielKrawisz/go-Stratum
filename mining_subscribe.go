package Stratum

type Subscription struct{}

type SubscribeRequestParams struct{}

type SubscribeResponseParams struct{}

type SubscribeRequest struct {
	Request
}

func (r *SubscribeRequest) Params() SubscribeRequestParams {}

type SubscribeResponse struct {
	Response
}

func (r *SubscribeResponse) Params() SubscribeResponseParams {}
