package response

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Json(w http.ResponseWriter, code int, v interface{}) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(bs)))
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)
	w.Write(bs)

	return nil
}
