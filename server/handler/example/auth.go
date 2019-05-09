package example

import "net/http"

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not_t"))
}
