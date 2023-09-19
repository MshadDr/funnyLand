package handler

import (
	"gitlab.com/M.darvish/funtory/internal/util/response"
	"net/http"
)

type HealthHandler struct {
	Code    int
	Message string
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health monitoring
func (h HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	h.Code = 222
	h.Message = "health is ok"
	_ = response.NewResponse(nil, h.Message, h.Code).Success(w)
	return
}
