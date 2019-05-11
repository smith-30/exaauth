package middleware

import (
	"context"
	"net/http"
	"time"

	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/smith-30/exaauth/logger"
	"go.uber.org/zap"
)

var (
	LoggerCtxKey = &contextKey{"Logger"}
)

// RequestLogger returns a logger handler using a custom LogFormatter.
func RequestLogger(zl *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			l := zl.With(zap.String("request_id", chi_middleware.GetReqID(r.Context())))
			ctx := r.Context()
			ctx = context.WithValue(ctx, LoggerCtxKey, l)

			// logging request info
			l.Info(
				"request started",
				zap.String("remote_addr", r.RemoteAddr),
				zap.Int64("req_size", r.ContentLength),
				zap.String("ua", r.UserAgent()),
				zap.String("path", r.URL.RequestURI()),
				zap.String("method", r.Method),
			)

			ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				l.Info("request finished", zap.Int("status", ww.Status()), zap.Int("body_size", ww.BytesWritten()), zap.Float64("duration_ms", float64(time.Since(t1))/float64(time.Millisecond)))
			}()

			next.ServeHTTP(ww, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetLogger(ctx context.Context) *logger.Logger {
	return ctx.Value(LoggerCtxKey).(*logger.Logger)
}
