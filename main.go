package main

import (
	"flag"
	"log/slog"
	"os"

	"git.bode.fun/meals/auth"
	mdb "git.bode.fun/meals/db"
	_ "modernc.org/sqlite"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := mainE(); err != nil {
		exitCode := 1
		slog.Error("program execution failed",
			"val", err.Error(),
			"exit_code", exitCode,
		)
		
		os.Exit(exitCode)
	}
}

func mainE() error {
	var addr string

	var dbPath string

	flag.StringVar(&addr, "addr", "127.0.0.1:3080", "The address to listen on")
	flag.StringVar(&dbPath, "db", "", "The path to the db")

	flag.Parse()

	cnf, err := ConfigFromEnv()
	if err != nil {
		return err
	}

	resolvedDBPath, err := mdb.ResolveDBPath(dbPath)
	if err != nil {
		return err
	}

	db, err := mdb.ConnectDB(resolvedDBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	err = mdb.Migrate(db)
	if err != nil {
		return err
	}

	auth, err := auth.New(auth.Config(cnf.Auth))
	if err != nil {
		return err
	}

	srv := NewServer(db, auth)

	slog.Info("starting server",
		"service", "server",
		"address", addr,
	)

	// TODO: Graceful shutdown
	return srv.ListenAndServe(addr)
}
