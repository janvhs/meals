package main

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/zitadel/oidc/v2/pkg/client/rs"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

// Errors
// -----------------------------------------------------------------------------

var (
	ErrCtxNotRegistered    = errors.New("auth: the auth provider is not registered on the context")
	ErrMissingAuthHeader   = errors.New("auth: request has no authorization header")
	ErrMissingBearerPrefix = errors.New("auth: the authorization header is missing the bearer prefix")
	ErrIntrospection       = errors.New("auth: token introspection failed")
	ErrUnauthenticated     = errors.New("auth: request is not authenticated")
)

// Context Keys
// -----------------------------------------------------------------------------

var providerKey = &contextKey{
	name: "auth: provider",
}

// Middleware
// -----------------------------------------------------------------------------

// TODO: Move provider creation on per request basis and handle errors via ctx?
func AuthMiddleware(
	issuer string,
	clientID string,
	keyID string,
	key string,
) (func(next http.Handler) http.Handler, error) {
	provider, err := newResourceServer(issuer, clientID, keyID, key)

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			r = registerProvider(provider, r)
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}, err
}

// Public Functions
// -----------------------------------------------------------------------------

// TODO: Check scopes?
func AuthenticateRequest(r *http.Request) (*oidc.IntrospectionResponse, error) {
	provider, ok := r.Context().Value(providerKey).(rs.ResourceServer)
	if !ok {
		return nil, ErrCtxNotRegistered
	}

	token, err := tokenFromRequest(r)
	if err != nil {
		return nil, err
	}

	resp, err := rs.Introspect(r.Context(), provider, token)
	if err != nil {
		return nil, ErrIntrospection
	}

	if !resp.Active {
		return resp, ErrUnauthenticated
	}

	return resp, err
}

// Private Functions
// -----------------------------------------------------------------------------

func newResourceServer(
	issuer string,
	clientID string,
	keyID string,
	key string,
) (rs.ResourceServer, error) {
	return rs.NewResourceServerJWTProfile(issuer, clientID, keyID, []byte(key))
}

func registerProvider(provider rs.ResourceServer, r *http.Request) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), providerKey, provider))
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
