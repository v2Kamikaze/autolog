package http

import (
	"net/http"
)

func ParseForm[T any](h *VehicleHandler, r *http.Request) (T, error) {
	var form T
	if err := r.ParseForm(); err != nil {
		return form, err
	}
	if err := h.formDecoder.Decode(&form, r.PostForm); err != nil {
		return form, err
	}
	return form, nil
}
