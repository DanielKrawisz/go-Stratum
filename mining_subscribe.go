package Stratum

// A Subscription is a 2-element json array containing a method
// and a session id.
type Subscription struct {
	Method Method
	ID     SessionID
}

type SubscribeRequestParams struct {
	UserAgent   string
	ExtraNonce1 *SessionID
}

type SubscribeResponseParams struct {
	Subscriptions []Subscription
	ExtraNonce1   SessionID
	uint32        ExtraNonce2Size
}

type subscribeRequest struct {
	Request
}

func (r *subscribeRequest) Params() (SubscribeRequestParams, error) {}

func NewSubscribeRequest(MessageID, SubscriptionRequestParams) *subscribeRequest {}

type subscribeResponse struct {
	Response
}

func (r *subscribeResponse) Params() (SubscribeResponseParams, error) {}

func NewSubscribeRequest(MessageID, SubscriptionResponseParams) *subscribeResponse {}
