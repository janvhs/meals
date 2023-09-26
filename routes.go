package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

func (s *Server) registerHandlers() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, http.StatusText(http.StatusOK))
	})

	s.Router.Get("/me", func(w http.ResponseWriter, r *http.Request) {
		auth, err := NewAuthService(AuthConfig(s.Cnf.Auth))
		if err != nil {
			slog.Error("initializing auth service", "service", "auth", "msg", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ir, err := auth.AuthenticateRequest(r)
		if err != nil {
			if errors.Is(err, ErrUnauthenticated) {
				slog.Warn("unauthenticated request", "service", "auth", "msg", err.Error())
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			} else if errors.Is(err, ErrIntrospection) {
				slog.Error("failed introspection", "service", "auth", "msg", err.Error())
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusInternalServerError)
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%#v\n\n%s\n%s\n", ir.Claims, ir.Subject, ir.Email)
	})
}
