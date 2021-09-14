package Stratum

type RequestID uint64

type Request struct {
	RequestID RequestID
	Method    Method
	params    [][]byte
}

type Response struct {
	RequestID RequestID
	result    []byte
	Error     *Error
}

type Notification struct {
	Method Method
	params [][]byte
}

func (r *Request) MarshallJSON() ([]byte, error) {}

func (r *Request) UnmarshallJSON([]byte) error {}

func (r *Response) MarshallJSON() ([]byte, error) {}

func (r *Response) UnmarshallJSON([]byte) error {}

func (n *Notification) MarshallJSON() ([]byte, error) {}

func (n *Notification) UnmarshallJSON([]byte) error {}
