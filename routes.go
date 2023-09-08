package main

import (
	"fmt"
	"net/http"
)

func (s *Server) registerHandlers() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, http.StatusText(http.StatusOK))
	})
}
