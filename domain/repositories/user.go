package repositories

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type UserRepository interface {
	GetUsersByQuery(ctx context.Context, q *entities.UsersQuery) ([]*entities.User, error)
	GetUsersBySerials(ctx context.Context, serials []string) ([]*entities.User, error)
	GetUserByID(ctx context.Context, id int64) (*entities.User, error)
	GetUserBySerial(ctx context.Context, serial string) (*entities.User, error)
	GetUserByEmail(ctx context.Context, name string) (*entities.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id int64) (*entities.User, error)
	DeleteUserBulk(ctx context.Context, ids []int64) ([]*entities.User, error)
	GetUserDetailsByUserSerials(ctx context.Context, userSerials []string) ([]*entities.UserDetail, error)
	GetMapUserDetailByUserSerial(ctx context.Context, userSerials []string) (map[string]*entities.UserDetail, error)
	UpsertUserDetail(ctx context.Context, userDetail *entities.UserDetail) error
}
