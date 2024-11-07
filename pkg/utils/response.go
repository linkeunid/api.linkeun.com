package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseOpts struct {
	Code        int         `json:"code"`
	Data        interface{} `json:"data"`
	Error       bool        `json:"error"`
	Message     string      `json:"message"`
	Limit       *int        `json:"limit,omitempty"`  // Optional field
	Offset      *int        `json:"offset,omitempty"` // Optional field
	Total       *int        `json:"total,omitempty"`  // Optional field
	TotalPages  *int        `json:"total_pages,omitempty"`
	CurrentPage *int        `json:"current_page,omitempty"`
	HasNext     *bool       `json:"has_next,omitempty"`
	HasPrev     *bool       `json:"has_prev,omitempty"`
}

func WriteJSONResponse(w http.ResponseWriter, opts *ResponseOpts) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(opts.Code)
	json.NewEncoder(w).Encode(opts)
}
