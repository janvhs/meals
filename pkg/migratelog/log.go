package migratelog

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pressly/goose/v3"
)

var _ goose.Logger = (*StructuredLogger)(nil)

type StructuredLogger struct{}

// Fatalf implements goose.Logger.
func (*StructuredLogger) Fatalf(format string, v ...interface{}) {
	slog.Error("migration failed",
		"service", "db",
		"val", fmt.Sprintf(format, v...),
	)
	os.Exit(1)
}

// Printf implements goose.Logger.
func (*StructuredLogger) Printf(format string, v ...interface{}) {
	slog.Info("migration",
		"service", "db",
		"val", fmt.Sprintf(format, v...),
	)
}
