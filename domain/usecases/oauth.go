package usecases

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type OAuthUsecase interface {
	CreateToken(ctx context.Context, r *entities.OAuthCreateTokenRequest) (*entities.OAuthCreateTokenResponse, error)
}
