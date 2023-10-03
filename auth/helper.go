package auth

import (
	"errors"
	"log/slog"
	"net/http"

	"git.bode.fun/meals/internal/httphelper"
)

func HandleError(err error, w http.ResponseWriter) {
	if errors.Is(err, errUnauthenticated) {
		slog.Warn("request unauthenticated", "service", "auth", "msg", err.Error())
		httphelper.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	} else if errors.Is(err, errIntrospection) {
		slog.Error("introspection failed", "service", "auth", "msg", err.Error())
		httphelper.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusInternalServerError)
	} else {
		httphelper.Error(w, err.Error(), http.StatusUnauthorized)
	}
}
