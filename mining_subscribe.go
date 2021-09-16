package Stratum

// A Subscription is a 2-element json array containing a method
// and a session id.
type Subscription struct {
	Method Method
	ID     SessionID
}

type SubscribeParams struct {
	UserAgent   string
	ExtraNonce1 *SessionID
}

func (p *SubscribeParams) Read(r *request) error {}

func SubscribeRequest(MessageID, SubscribeParams) request {}

type SubscribeResult struct {
	Subscriptions   []Subscription
	ExtraNonce1     SessionID
	ExtraNonce2Size uint32
}

func (p *SubscribeResult) Read(r *response) error {}

func SubscribeResponse(MessageID, SubscribeResult) response {}
