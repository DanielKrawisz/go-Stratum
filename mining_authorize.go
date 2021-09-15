package Stratum

type AuthorizeParams struct {
	Username string

	// Password is optional. Pools don't necessarily require a miner to log in to mine.
	Password *string
}

type authorizeRequest struct {
	Request
}

func (au *authorizeRequest) Params() (AuthorizeParams, error) {}

func NewAuthorizeRequest(RequestID, AuthorizeParams) *authorizeRequest {}

type authorizeResponse booleanResponse
