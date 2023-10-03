package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"git.bode.fun/meals/auth"
	"git.bode.fun/meals/db/user"
	"git.bode.fun/meals/internal/httphelper"
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

		userRepo := user.New(s.DB)
		err = userRepo.EnsureExists(r.Context(), ir.Subject)
		if err != nil {
			slog.Error("ensuring user reference", "service", "user repo", "msg", err.Error())
			httphelper.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		user, err := userRepo.Get(r.Context(), ir.Subject)
		if err != nil {
			slog.Error("get user reference", "service", "user repo", "msg", err.Error())
			httphelper.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", user.ID)
	})
}
