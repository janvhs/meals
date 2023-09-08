package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router chi.Router
}

func NewServer() *Server {
	r := chi.NewRouter()

	srv := &Server{
		Router: r,
	}

	srv.registerMiddleware()

	srv.registerHandlers()

	return srv
}

// Public Methods
// ------------------------------------------------------------------------

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Router.ServeHTTP(w, req)
}

func (s *Server) ListenAndServe(addr string) error {
	srv := new(http.Server)
	srv.Addr = addr
	srv.Handler = s
	srv.ReadHeaderTimeout = time.Duration(time.Duration(10).Seconds())
	srv.WriteTimeout = time.Duration(time.Duration(10).Seconds())

	return srv.ListenAndServe()
}

// Private Methods
// ------------------------------------------------------------------------

func (s *Server) registerMiddleware() {
	// Give each request a unique ID
	s.Router.Use(middleware.RequestID)
	// Get the client's ip even when proxied
	s.Router.Use(middleware.RealIP)
	// Remove multiple slashes from the requested resource path
	s.Router.Use(middleware.CleanPath)
	// Remove any trailing slash from the requested resource path
	s.Router.Use(middleware.StripSlashes)

	// TODO: Add some kind of rate limiting (uber)
	// TODO: Add CORS (rs/cors) https://github.com/rs/cors/blob/master/examples/chi/server.go
	// TODO: Add CSRF protection
	// TODO: Add security headers (unrolled/secure)
	// TODO: Add compression
	// TODO: Add request timeout
	// TODO: use slog instead

	// Log every incoming request
	// Log middleware depends on Recover
	s.Router.Use(middleware.Logger)
	// A panic should not quit the program
	s.Router.Use(middleware.Recoverer)
}
