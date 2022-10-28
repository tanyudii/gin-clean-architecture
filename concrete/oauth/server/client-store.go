package oauth

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/vodeacloud/hr-api/domain/repositories"
)

type clientStore struct {
	clientRepo repositories.OAuthClientRepository
}

func newClientStore(clientRepo repositories.OAuthClientRepository) oauth2.ClientStore {
	return &clientStore{
		clientRepo: clientRepo,
	}
}

func (s *clientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	return s.clientRepo.GetOAuthClientByClientID(ctx, id)
}
