package services

import (
	"context"
	"log/slog"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type UserClaims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type AuthService struct {
	Logger    *slog.Logger
	Oauth2Cfg *oauth2.Config
	OidcP     *oidc.Provider
}

func NewAuthService(
	logger *slog.Logger,
	oauth2 *oauth2.Config,
	oidcP *oidc.Provider,
) *AuthService {
	return &AuthService{
		Logger:    logger.WithGroup("Auth Service"),
		Oauth2Cfg: oauth2,
		OidcP:     oidcP,
	}
}

func (s *AuthService) GenerateAuthUrl() string {
	return s.Oauth2Cfg.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// this will return the raw jwt
func (s *AuthService) Callback(ctx context.Context, queryCode string) (string, error) {
	oauth2Token, err := s.Oauth2Cfg.Exchange(ctx, queryCode)
	if err != nil {
		s.Logger.Error("Failed to exchange token", slog.String("error", err.Error()))
		return "", err
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		s.Logger.Error("No id_token field in oauth2 token")
		return "", err
	}

	return rawIDToken, nil
}

// get the token claims
func (s *AuthService) GetTokenClaims(
	ctx context.Context,
	rawIDToken string,
) (UserClaims, error) {
	idToken, err := s.OidcP.Verifier(&oidc.Config{ClientID: s.Oauth2Cfg.ClientID}).Verify(ctx, rawIDToken)
	if err != nil {
		s.Logger.Error("Failed to verify ID Token", slog.String("error", err.Error()))
		return UserClaims{}, err
	}

	var claims UserClaims

	if err := idToken.Claims(&claims); err != nil {
		return UserClaims{}, err
	}

	return claims, nil
}
