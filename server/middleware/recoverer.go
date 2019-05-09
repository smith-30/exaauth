package middleware

// I made this with reference to github.com/go-chi/chi/middleware/recoverer.go .
// The reason is that I wanted to use zap for logging.

import (
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"
)

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
//
// Alternatively, look at https://github.com/pressly/lg middleware pkgs.
func Recoverer() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					GetLogger(r.Context()).Error("", zap.Any("panic_by", rvr), zap.String("stack", string(debug.Stack())))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
