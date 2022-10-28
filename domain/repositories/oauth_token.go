package repositories

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type OAuthTokenRepository interface {
	CreateAccessToken(ctx context.Context, at *entities.AccessToken) error
	CreateRefreshToken(ctx context.Context, rt *entities.RefreshToken) error
	GetAccessToken(ctx context.Context, id string) (*entities.AccessToken, error)
	GetRefreshToken(ctx context.Context, id string) (*entities.RefreshToken, error)
	DeleteAccessToken(ctx context.Context, id string) error
	DeleteRefreshToken(ctx context.Context, id string) error
}
