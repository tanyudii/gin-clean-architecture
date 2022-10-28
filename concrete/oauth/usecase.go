package oauth

import (
	"context"
	"errors"
	"github.com/go-oauth2/oauth2/v4"
	oautherr "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"
	oauth "github.com/vodeacloud/hr-api/concrete/oauth/server"
	"github.com/vodeacloud/hr-api/config"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
	"time"
)

type Usecase struct {
	clientRepo repositories.OAuthClientRepository
	userUc     usecases.UserUsecase
	roleRepo   repositories.RoleRepository
	oauth2Svr  *server.Server
}

func NewUsecase(
	cfg *config.Config,
	tokenRepo repositories.OAuthTokenRepository,
	clientRepo repositories.OAuthClientRepository,
	userUc usecases.UserUsecase,
	roleRepo repositories.RoleRepository,
) usecases.OAuthUsecase {
	return &Usecase{
		clientRepo: clientRepo,
		userUc:     userUc,
		roleRepo:   roleRepo,
		oauth2Svr:  oauth.NewOAuth2Server(cfg, clientRepo, tokenRepo),
	}
}

func (u *Usecase) CreateToken(ctx context.Context, r *entities.OAuthCreateTokenRequest) (resp *entities.OAuthCreateTokenResponse, err error) {
	if err = r.Validate(); err != nil {
		return
	}

	tokenGenerateReq := &oauth2.TokenGenerateRequest{
		ClientID:     r.ClientID,
		ClientSecret: r.ClientSecret,
		Scope:        r.Scope,
	}

	if r.GrantType == "password" {
		if err = u.grantTypePassword(ctx, r, tokenGenerateReq); err != nil {
			return
		}
	} else if r.GrantType == "refresh_token" {
		tokenGenerateReq.Refresh = r.Refresh
	}

	tokenInfo, err := u.oauth2Svr.GetAccessToken(ctx, oauth2.GrantType(r.GrantType), tokenGenerateReq)
	if err != nil {
		if errors.Is(err, oautherr.ErrInvalidClient) {
			err = entities.ErrInvalidClient
		}
		return nil, err
	}

	expiresAt := tokenInfo.GetAccessCreateAt().Add(tokenInfo.GetAccessExpiresIn())
	return &entities.OAuthCreateTokenResponse{
		AccessToken:  tokenInfo.GetAccess(),
		RefreshToken: tokenInfo.GetRefresh(),
		ExpiresIn:    int32(expiresAt.Sub(time.Now()).Seconds()),
		TokenType:    "Bearer",
		Scope:        tokenInfo.GetScope(),
	}, nil
}
