package auth

import (
	"context"
	"errors"
	"fmt"
	"mergedup/business/core/user"
	"mergedup/business/core/user/repositories/userdb"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// AuthError is used to pass an error during the request through the
// application with auth specific context.
type AuthError struct {
	msg string
}

// NewAuthError creates an AuthError for the provided message.
func NewAuthError(format string, args ...any) error {
	return &AuthError{
		msg: fmt.Sprintf(format, args...),
	}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (ae *AuthError) Error() string {
	return ae.msg
}

// IsAuthError checks if an error of type AuthError exists.
func IsAuthError(err error) bool {
	var ae *AuthError
	return errors.As(err, &ae)
}

// These the current set of rules we have for auth.
const (
	RuleAuthenticate = "auth"
	RuleAny          = "allowAny"
	RuleAdminOnly    = "allowOnlyAdmin"
	RuleUserOnly     = "allowOnlyUser"
	SecretKey        = "secret"
)

// ErrForbidden is returned when a auth issue is identified.
var ErrForbidden = errors.New("attempted action is not allowed")

// Config represents information required to initialize auth.
type Config struct {
	Log *zap.SugaredLogger
	DB  *sqlx.DB
}

// Auth is used to authenticate clients. It can generate a token for a
// set of user claims and recreate the claims by parsing the token.
type Auth struct {
	log    *zap.SugaredLogger
	user   *user.Core
	method jwt.SigningMethod
	parser *jwt.Parser
	cache  map[string]string
}

// New creates an Auth to support authentication/authorization.
func New(cfg Config) (*Auth, error) {

	// If a database connection is not provided, we won't perform the
	// user enabled check.
	var usr *user.Core
	if cfg.DB != nil {
		usr = user.NewCore(userdb.NewStore(cfg.Log, cfg.DB))
	}

	a := Auth{
		log:    cfg.Log,
		user:   usr,
		method: jwt.GetSigningMethod(jwt.SigningMethodRS256.Name),
		parser: jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name})),
		cache:  make(map[string]string),
	}

	return &a, nil
}

// Authenticate processes the token to validate the sender's token is valid.
func (a *Auth) Authenticate(ctx context.Context, bearerToken string) (Claims, error) {
	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return Claims{}, errors.New("expected authorization header format: Bearer <token>")
	}

	var claims Claims

	token, err := jwt.ParseWithClaims(parts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Roles, claims.RegisteredClaims.Issuer)
	} else {
		fmt.Println(err)
	}

	// Check the database for this user to verify they are still enabled.
	if !a.isUserEnabled(ctx, claims) {
		return Claims{}, fmt.Errorf("user not enabled : %w", err)
	}

	return claims, nil
}

// Authorize attempts to authorize the user with the provided input roles, if
// none of the input roles are within the user's claims, we return an error
// otherwise the user is authorized.
func (a *Auth) Authorize(ctx context.Context, claims Claims, rule string) error {

	if rule == RuleAny {
		for _, role := range claims.Roles {
			if role == user.RoleAdmin || role == user.RoleUser {
				return nil
			}
		}
	} else if rule == RuleAdminOnly {
		for _, role := range claims.Roles {
			if role == user.RoleAdmin {
				return nil
			}
		}

	} else if rule == RuleUserOnly {
		for _, role := range claims.Roles {
			if role == user.RoleUser {
				return nil
			}
		}
	}

	return ErrForbidden
}

// isUserEnabled hits the database and checks the user is not disabled. If the
// no database connection was provided, this check is skipped.
func (a *Auth) isUserEnabled(ctx context.Context, claims Claims) bool {
	if a.user == nil {
		return true
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return false
	}

	usr, err := a.user.QueryByID(ctx, userID)
	if err != nil {
		return false
	}

	return usr.Active
}

// GenerateToken generates a signed JWT token string representing the user Claims.
func (a *Auth) GenerateToken(kid string, claims Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)

	str, err := token.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return str, nil
}
