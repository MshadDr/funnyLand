package util

import (
	"encoding/json"
	"net/http"
)

const (
	SuccessStatus = "Success"
	FailedStatus  = "Failed"
	SuccessCode   = 200
	FailedCode    = 400
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponse the Response constructor
func NewResponse(data interface{}, message string, code int) *Response {
	return &Response{
		Status:  SuccessStatus,
		Message: message,
		Code:    code,
		Data:    data,
	}
}

// set json header
func setJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
}

// Success returns a success response with data, message, and status code.
func (r *Response) Success(w http.ResponseWriter) error {
	statusCode := r.Code
	if r.Code == 0 {
		statusCode = SuccessCode
	}

	setJSONHeaders(w)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		return err
	}
	return nil
}

// Failed returns a failed response with data, message, and status code.
func (r *Response) Failed(w http.ResponseWriter) error {
	r.Status = FailedStatus
	statusCode := r.Code
	if r.Code == 0 {
		statusCode = FailedCode
	}

	setJSONHeaders(w)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		return err
	}
	return nil
}
