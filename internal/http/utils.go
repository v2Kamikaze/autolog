package http

import (
	"net/http"
	"strconv"
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

func ParseInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}

	return i
}
