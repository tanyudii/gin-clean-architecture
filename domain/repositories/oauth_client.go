package repositories

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type OAuthClientRepository interface {
	GetOAuthClientsByQuery(ctx context.Context, q *entities.OAuthClientsQuery) ([]*entities.OAuthClient, error)
	GetOAuthClientByID(ctx context.Context, id int64) (*entities.OAuthClient, error)
	GetOAuthClientByName(ctx context.Context, name string) (*entities.OAuthClient, error)
	GetOAuthClientByClientID(ctx context.Context, clientID string) (*entities.OAuthClient, error)
	CreateOAuthClient(ctx context.Context, role *entities.OAuthClient) error
	UpdateOAuthClient(ctx context.Context, role *entities.OAuthClient) error
	DeleteOAuthClient(ctx context.Context, id int64) (*entities.OAuthClient, error)
	DeleteOAuthClientBulk(ctx context.Context, ids []int64) ([]*entities.OAuthClient, error)
}
