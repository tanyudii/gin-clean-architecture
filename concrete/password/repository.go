package password

import (
	"github.com/go-redis/redis/v8"
	"github.com/vodeacloud/hr-api/domain/repositories"
)

type Repository struct {
	redisCli *redis.Client
}

func NewRepository(
	redisCli *redis.Client,
) repositories.PasswordRepository {
	return &Repository{
		redisCli: redisCli,
	}
}
