package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"git.bode.fun/meals/auth"
)

func (s *Server) registerHandlers() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, http.StatusText(http.StatusOK))
	})

	s.Router.Get("/me", func(w http.ResponseWriter, r *http.Request) {
		ir, err := s.Auth.AuthenticateRequest(r)
		if err != nil {
			auth.HandleError(err, w)

			return
		}

		userRepo := NewUserRepository(s.DB)
		err = userRepo.EnsureExists(ir.Subject)
		if err != nil {
			slog.Error("ensuring user reference", "service", "user repo", "msg", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		user, err := userRepo.UserByID(ir.Subject)
		if err != nil {
			slog.Error("get user reference", "service", "user repo", "msg", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", user.ID)
	})
}
