package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/zitadel/oidc/v2/pkg/client/rs"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

// Errors
// -----------------------------------------------------------------------------

var (
	errMissingAuthHeader   = errors.New("auth: request has no authorization header")
	errMissingBearerPrefix = errors.New("auth: the authorization header is missing the bearer prefix")
	errIntrospection       = errors.New("auth: token introspection failed")
	errUnauthenticated     = errors.New("auth: request is not authenticated")
)

// Service
// -----------------------------------------------------------------------------

type Service interface {
	AuthenticateRequest(r *http.Request) (*oidc.IntrospectionResponse, error)
}

type Config struct {
	Issuer   string
	ClientID string
	KeyID    string
	Key      string
}

var _ Service = (*Authenticator)(nil)

type Authenticator struct {
	provider rs.ResourceServer
}

func New(
	cnf Config,
) (*Authenticator, error) {
	provider, err := newResourceServer(cnf.Issuer, cnf.ClientID, cnf.KeyID, cnf.Key)
	if err != nil {
		return nil, err
	}

	return &Authenticator{
		provider: provider,
	}, nil
}

// Public Methods
// -----------------------------------------------------------------------------

func (a Authenticator) AuthenticateRequest(r *http.Request) (*oidc.IntrospectionResponse, error) {
	token, err := tokenFromRequest(r)
	if err != nil {
		return nil, err
	}

	resp, err := rs.Introspect(r.Context(), a.provider, token)
	if err != nil {
		return nil, errIntrospection
	}

	if !resp.Active {
		return resp, errUnauthenticated
	}

	// TODO: Check scopes?

	return resp, nil
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
		return "", errMissingAuthHeader
	}

	if !strings.HasPrefix(auth, oidc.PrefixBearer) {
		return "", errMissingBearerPrefix
	}

	return strings.TrimPrefix(auth, oidc.PrefixBearer), nil
}
