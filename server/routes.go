package server

import (
	"context"
	"errors"
	"net/http"

	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/smith-30/petit/logger"
	"github.com/smith-30/petit/server/handler/example"
	"github.com/smith-30/petit/server/middleware"
	"go.uber.org/zap"
)

type Server struct {
	server *http.Server
	Logger *logger.Logger

	// serverに連携先を登録できるようにしたい

	// RDB はServerでいいかな？
}

func NewServer(options ...func(*Server)) *Server {
	a := &Server{}
	for _, option := range options {
		option(a)
	}
	return a
}

func Logger(zl *logger.Logger) func(*Server) {
	return func(a *Server) {
		a.Logger = &logger.Logger{zl.Named("Server")}
	}
}

func Address(host, port string) func(*Server) {
	return func(a *Server) {
		s := &http.Server{
			Addr: host + ":" + port,
		}
		a.server = s
	}
}

func (a *Server) Start() error {
	a.Logger.Info("server started.", zap.String("addr", a.server.Addr))
	return a.server.ListenAndServe()
}

func (a *Server) Shutdown(ctx context.Context) error {
	a.Logger.Info("server stopped.")
	return a.server.Shutdown(ctx)
}

func Routes() func(*Server) {
	return func(a *Server) {
		r := chi.NewRouter()

		// Basic CORS
		// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
		cors := cors.New(cors.Options{
			// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		r.Use(cors.Handler)

		//
		// middlewares
		// chi: all middlewares must be defined before routes on a mux
		//
		r.Use(chi_middleware.RequestID)
		r.Use(chi_middleware.RealIP)

		// link logger to context of request
		r.Use(middleware.RequestLogger(a.Logger))

		r.Use(middleware.Recoverer())

		//
		// your application routing...
		//
		r.Route("/api", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				panic(errors.New("test"))
			})
			r.Get("/auth", example.Auth)
		})

		// for debug
		r.Route("/debug", func(r chi.Router) {
			r.Get("/stats", stats_api.Handler)
		})

		a.server.Handler = r
	}
}
