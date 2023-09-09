package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	if err := mainE(); err != nil {
		exitCode := 1
		slog.Error(err.Error(), "exit_code", exitCode)
		os.Exit(exitCode)
	}
}

func mainE() error {
	var addr string

	var dbPath string

	flag.StringVar(&addr, "addr", "127.0.0.1:3080", "The address to listen on")
	flag.StringVar(&dbPath, "db", "", "The path to the db")

	flag.Parse()

	resolvedDBPath, err := ResolveDBPath(dbPath)
	if err != nil {
		return err
	}

	db, err := ConnectDB(resolvedDBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	err = Migrate(db.DB)
	if err != nil {
		return err
	}

	srv := NewServer(db)

	slog.Info("Server is starting", "address", addr)

	// TODO: Graceful shutdown
	return srv.ListenAndServe(addr)
}
