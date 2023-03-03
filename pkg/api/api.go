package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/b1izko/test-binance-api/pkg/api/errors"
	"github.com/davecgh/go-spew/spew"
)

// Response model
type Response struct {
	Success bool          `json:"success"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *errors.Error `json:"error,omitempty"`
}

// NewResponse returns new response
func NewResponse(result interface{}) *Response {
	return &Response{Success: true, Result: result}
}

// WriteResponse to response writer
func (r *Response) WriteResponse(w http.ResponseWriter) {
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("MarshalResponse: %v\n", spew.Sdump(r.Result))
		fmt.Println("MarshalResponse", err)
		data = []byte(`{"success":false,"error":{"error": 10, "error_msg":"cannot marshal response, see server logs for details"}}`)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

// WriteError create error response
func WriteError(w http.ResponseWriter, code int, msg string) {
	err := &errors.Error{Code: code, Msg: msg}
	resp := Response{
		Error: err,
	}
	resp.WriteResponse(w)
}
