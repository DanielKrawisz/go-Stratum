package Stratum

type AuthorizeParams struct {
	Username string
	Password *string
}

type authorizeRequest struct {
	Request
}

func (au *authorizeRequest) Params() AuthorizeParams {}

func NewAuthorizeRequest(RequestID, AuthorizeParams) (*authorizeRequest, error) {}

type authorizeResponse BooleanResponse
