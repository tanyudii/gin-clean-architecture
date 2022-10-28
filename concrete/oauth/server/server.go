package oauth

import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/vodeacloud/hr-api/config"
	"github.com/vodeacloud/hr-api/domain/repositories"
)

func NewOAuth2Server(
	cfg *config.Config,
	clientRepo repositories.OAuthClientRepository,
	tokenRepo repositories.OAuthTokenRepository,
) *server.Server {
	clientStorage := newClientStore(clientRepo)
	accessGen := newAccessGenerate(cfg)
	tokenStorage := newTokenStore(tokenRepo, accessGen)
	manager := manage.NewDefaultManager()
	manager.MapClientStorage(clientStorage)
	manager.MapAccessGenerate(accessGen)
	manager.MapTokenStorage(tokenStorage)
	return server.NewDefaultServer(manager)
}
