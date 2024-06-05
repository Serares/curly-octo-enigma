package services

import (
	"context"
	"log/slog"

	"github.com/Serares/curly-octo-enigma/api/muxIntegration/utils"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type CredentialsService struct {
	Logger    *slog.Logger
	Oauth2Cfg *oauth2.Config
	OidcP     *oidc.Provider
}

func NewCredentialsService(
	logger *slog.Logger,
	oauth2Cfg *oauth2.Config,
	oidcP *oidc.Provider,
) *CredentialsService {
	return &CredentialsService{
		Logger:    logger.WithGroup("MuxCredentials Service"),
		Oauth2Cfg: oauth2Cfg,
		OidcP:     oidcP,
	}
}

// get the token claims
func (s *CredentialsService) GetTokenClaims(
	ctx context.Context,
	rawIDToken string,
) (utils.UserClaims, error) {
	idToken, err := s.OidcP.Verifier(&oidc.Config{ClientID: s.Oauth2Cfg.ClientID}).Verify(ctx, rawIDToken)
	if err != nil {
		s.Logger.Error("Failed to verify ID Token", slog.String("error", err.Error()))
		return utils.UserClaims{}, err
	}

	var claims utils.UserClaims

	if err := idToken.Claims(&claims); err != nil {
		return utils.UserClaims{}, err
	}

	return claims, nil
}
