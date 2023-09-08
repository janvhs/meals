package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

var _ http.Handler = (*Server)(nil)

func main() {
	if err := mainE(); err != nil {
		exitCode := 1
		slog.Error(err.Error(), "exit_code", exitCode)
		os.Exit(exitCode)
	}
}

func mainE() error {
	var addr string

	flag.StringVar(&addr, "addr", "127.0.0.1:3080", "The address to listen on")

	flag.Parse()

	srv := NewServer()

	slog.Info("Starting server", "address", addr)

	return srv.ListenAndServe(addr)
}
