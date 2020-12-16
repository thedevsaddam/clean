package http

import (
	"encoding/json"
	"net/http"
)

//Response ...
type Response struct {
	Status  int         `json:"-"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

//Write ...
func (r *Response) JSON(w http.ResponseWriter) error {
	bb, err := json.Marshal(r)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Status != 0 {
		w.WriteHeader(r.Status)
	}
	_, err = w.Write(bb)
	return err
}
