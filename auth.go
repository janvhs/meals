package main

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/zitadel/oidc/v2/pkg/client/rs"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

// Errors
// -----------------------------------------------------------------------------

var (
	ErrMissingAuthHeader   = errors.New("auth: request has no authorization header")
	ErrMissingBearerPrefix = errors.New("auth: the authorization header is missing the bearer prefix")
	ErrIntrospection       = errors.New("auth: token introspection failed")
	ErrUnauthenticated     = errors.New("auth: request is not authenticated")
)

// Service
// -----------------------------------------------------------------------------

type AuthConfig struct {
	Issuer   string
	ClientID string
	KeyID    string
	Key      string
}

type AuthService struct {
	provider rs.ResourceServer
}

func NewAuthService(
	cnf AuthConfig,
) (*AuthService, error) {
	provider, err := newResourceServer(cnf.Issuer, cnf.ClientID, cnf.KeyID, cnf.Key)
	if err != nil {
		return nil, err
	}

	return &AuthService{
		provider: provider,
	}, nil
}

// Public Methods
// -----------------------------------------------------------------------------

func (a *AuthService) AuthenticateRequest(r *http.Request) (*oidc.IntrospectionResponse, error) {
	token, err := tokenFromRequest(r)
	if err != nil {
		return nil, err
	}

	resp, err := rs.Introspect(r.Context(), a.provider, token)
	if err != nil {
		return nil, ErrIntrospection
	}

	if !resp.Active {
		return resp, ErrUnauthenticated
	}

	// TODO: Check scopes?

	return resp, nil
}

// Public Functions

func HandleAuthError(err error, w http.ResponseWriter) {
	if errors.Is(err, ErrUnauthenticated) {
		slog.Warn("unauthenticated request", "service", "auth", "msg", err.Error())
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	} else if errors.Is(err, ErrIntrospection) {
		slog.Error("failed introspection", "service", "auth", "msg", err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusInternalServerError)
	} else {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
}

// Private Functions
// -----------------------------------------------------------------------------

func newResourceServer( //nolint
	issuer string,
	clientID string,
	keyID string,
	key string,
) (rs.ResourceServer, error) {
	return rs.NewResourceServerJWTProfile(issuer, clientID, keyID, []byte(key))
}

func tokenFromRequest(r *http.Request) (string, error) {
	auth := r.Header.Get("authorization")
	if auth == "" {
		return "", ErrMissingAuthHeader
	}

	if !strings.HasPrefix(auth, oidc.PrefixBearer) {
		return "", ErrMissingBearerPrefix
	}

	return strings.TrimPrefix(auth, oidc.PrefixBearer), nil
}
