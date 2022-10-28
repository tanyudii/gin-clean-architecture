package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"time"
)

type TokenRepository struct {
	redisClient *redis.Client
}

func NewTokenRepository(
	redisClient *redis.Client,
) repositories.OAuthTokenRepository {
	return &TokenRepository{
		redisClient: redisClient,
	}
}

func (r *TokenRepository) CreateAccessToken(ctx context.Context, at *entities.AccessToken) error {
	jsonAt, err := json.Marshal(at)
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = r.redisClient.Set(ctx, keyAccessToken(at.ID), jsonAt, at.ExpiresAt.Sub(now)).Result()
	return err
}

func (r *TokenRepository) CreateRefreshToken(ctx context.Context, rt *entities.RefreshToken) error {
	jsonAt, err := json.Marshal(rt)
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = r.redisClient.Set(ctx, keyRefreshToken(rt.ID), jsonAt, rt.ExpiresAt.Sub(now)).Result()
	return err
}

func (r *TokenRepository) GetAccessToken(ctx context.Context, id string) (*entities.AccessToken, error) {
	data, err := r.redisClient.Get(ctx, keyAccessToken(id)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = errutil.NewNotFoundError("access token not found")
		}
		return nil, err
	}
	at := new(entities.AccessToken)
	if err = json.Unmarshal([]byte(data), at); err != nil {
		return nil, err
	}
	return at, nil
}

func (r *TokenRepository) GetRefreshToken(ctx context.Context, id string) (*entities.RefreshToken, error) {
	refreshKey := keyRefreshToken(id)
	data, err := r.redisClient.Get(ctx, refreshKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = errutil.NewNotFoundError("refresh token not found")
		}
		return nil, err
	}
	rt := new(entities.RefreshToken)
	if err = json.Unmarshal([]byte(data), rt); err != nil {
		return nil, err
	}
	return rt, nil
}

func (r *TokenRepository) DeleteAccessToken(ctx context.Context, id string) error {
	return r.redisClient.Del(ctx, keyAccessToken(id)).Err()
}

func (r *TokenRepository) DeleteRefreshToken(ctx context.Context, id string) error {
	return r.redisClient.Del(ctx, keyRefreshToken(id)).Err()
}
