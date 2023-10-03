package httplog

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

var _ middleware.LogFormatter = (*StructuredLogger)(nil)

type StructuredLogger struct {
	Logger *slog.Logger
}

// NewLogEntry implements middleware.LogFormatter.
func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry { //nolint
	var logAttrs []slog.Attr

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logAttrs = append(logAttrs, slog.String("req_id", reqID))
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	handler := l.Logger.Handler().WithAttrs(append(logAttrs,
		slog.String("http_scheme", scheme),
		slog.String("http_proto", r.Proto),
		slog.String("http_method", r.Method),
		slog.String("remote_addr", r.RemoteAddr),
		slog.String("user_agent", r.UserAgent()),
		slog.String("uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)),
	))

	entry := StructuredLogEntry{
		Logger: slog.New(handler),
	}

	entry.Logger.Info("request started")

	return &entry
}

var _ middleware.LogEntry = (*StructuredLogEntry)(nil)

type StructuredLogEntry struct {
	Logger *slog.Logger
}

// Panic implements middleware.LogEntry.
func (e *StructuredLogEntry) Panic(v interface{}, stack []byte) {
	e.Logger.Error("request panicked",
		"stack", string(stack), "panic",
		fmt.Sprintf("%+v", v),
	)
}

// Write implements middleware.LogEntry.
func (e *StructuredLogEntry) Write(status int, bytes int, _ http.Header, elapsed time.Duration, _ interface{}) {
	e.Logger.Info("request complete",
		"resp_status", status,
		"resp_byte_length", bytes,
		"resp_elapsed_ms", elapsed.Microseconds(),
	)
}
