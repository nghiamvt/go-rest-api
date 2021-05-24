package jwt

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"go.uber.org/zap"

	"github.com/0xTatsu/mvtn-api/config"
)

type AuthJWT struct {
	cfg     *config.Configuration
	jwtAuth *jwtauth.JWTAuth
}

func NewJWT(
	cfg *config.Configuration,
) *AuthJWT {
	return &AuthJWT{
		cfg:     cfg,
		jwtAuth: jwtauth.New("HS256", []byte(cfg.JWT.Secret), nil),
	}
}

// CreateAccessToken returns an access token for provided account claims.
func (a *AuthJWT) CreateAccessToken(c AccessClaims) (string, error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Hour * time.Duration(a.cfg.JWT.ExpiryInHour)).Unix()

	_, tokenString, err := a.jwtAuth.Encode(ToMapStringInterface(c))

	return tokenString, err
}

// CreateRefreshToken returns a refresh token for provided token Claims.
func (a *AuthJWT) CreateRefreshToken(c RefreshClaims) (string, error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Hour * time.Duration(a.cfg.JWT.RefreshExpiryInHour)).Unix()

	_, tokenString, err := a.jwtAuth.Encode(ToMapStringInterface(c))

	return tokenString, err
}

// CreateTokenPair returns both an access token and a refresh token.
func (a *AuthJWT) CreateTokenPair(accessClaims AccessClaims, refreshClaims RefreshClaims) (string, string, error) {
	accessToken, err := a.CreateAccessToken(accessClaims)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := a.CreateRefreshToken(refreshClaims)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Verifier http middleware will verify a jwt string from a http request.
func (a *AuthJWT) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(a.jwtAuth)
}

func ToMapStringInterface(c interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	j, err := json.Marshal(c)
	if err != nil {
		zap.L().Error("cannot marshal", zap.Error(err))
	}

	if err := json.Unmarshal(j, &m); err != nil {
		zap.L().Error("cannot marshal", zap.Error(err))
	}

	return m
}