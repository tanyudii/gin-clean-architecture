package usecases

import (
	"context"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type UserUsecase interface {
	GetUsersByQuery(ctx context.Context, r *entities.GetUsersByQueryRequest) (*entities.UsersResponse, error)
	GetUserByID(ctx context.Context, r *entities.GetUserByIDRequest) (*entities.UserResponse, error)
	CreateUser(ctx context.Context, r *entities.CreateUserRequest) (*entities.UserResponse, error)
	UpdateUser(ctx context.Context, r *entities.UpdateUserRequest) (*entities.UserResponse, error)
	DeleteUser(ctx context.Context, r *entities.DeleteUserRequest) (*entities.UserResponse, error)
	DeleteUserBulk(ctx context.Context, r *entities.DeleteUserBulkRequest) (*entities.UsersResponseWithoutPagination, error)
}
