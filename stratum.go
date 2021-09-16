package Stratum

import (
	"encoding/json"
	"errors"
)

// Stratum has three types of messages: notification, request, and response.

// Notification for methods that do not require a response.
type notification struct {
	method method        `json:"method"`
	params []interface{} `json:"params"`
}

func Notification(m Method, params []interface{}) notification {
	n, _ := EncodeMethod(m)
	return notification{
		method: n,
		params: params,
	}
}

func (n *notification) Method() Method {
	m, _ := DecodeMethod(n.method)
	return m
}

// request is for methods that require a response.
type request struct {
	MessageID MessageID     `json:"id"`
	method    method        `json:"method"`
	params    []interface{} `json:"params"`
}

func Request(id MessageID, m Method, params []interface{}) request {
	n, _ := EncodeMethod(m)
	return response{
		MessageID: id,
		method:    n,
		params:    params,
	}
}

func (n *request) Method() Method {
	m, _ := DecodeMethod(n.method)
	return m
}

// Response is what is sent back in response to requests.
type response struct {
	MessageID MessageID   `json:"id"`
	result    interface{} `json:"result"`
	Error     *Error      `json:"error"`
}

func Response(id MessageID, r interface{}) response {
	return response{
		MessageID: id,
		result:    r,
		Error:     nil,
	}
}

func ErrorResponse(id MessageID, e Error) response {
	return response{
		MessageID: id,
		result:    nil,
		Error:     &e,
	}
}

func (r *request) MarshallJSON() ([]byte, error) {
	if !ValidMessageID(r.MessageID) {
		return []byte{}, errors.New("Invalid id")
	}

	if r.method == "" {
		return []byte{}, errors.New("Invalid method")
	}

	return json.Marshal(r)
}

func (r *request) UnmarshallJSON(j []byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if !ValidMessageID(r.MessageID) {
		return errors.New("Invalid id")
	}

	if r.Method() == Unset {
		return errors.New("Invalid method")
	}
}

func (r *response) MarshallJSON() ([]byte, error) {
	if !ValidMessageID(r.MessageID) {
		return []byte{}, errors.New("Invalid id")
	}

	return json.Marshal(r)
}

func (r *response) UnmarshallJSON([]byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if !ValidMessageID(r.MessageID) {
		return errors.New("Invalid id")
	}
}

func (r *notification) MarshallJSON() ([]byte, error) {
	if r.method == "" {
		return []byte{}, errors.New("Invalid method")
	}

	return json.Marshal(r)
}

func (r *notification) UnmarshallJSON([]byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if r.Method() == Unset {
		return errors.New("Invalid method")
	}
}
