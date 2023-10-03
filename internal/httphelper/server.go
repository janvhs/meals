package httphelper

import (
	"log/slog"
	"net/http"

	"github.com/go-json-experiment/json"
)

type errResp struct {
	Error string
}

func Error(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	resp := errResp{
		Error: err,
	}

	if err := json.MarshalWrite(w, &resp); err != nil {
		slog.Error("response failed",
			"service", "httphelper",
			"val", err.Error(),
		)
	}
}
