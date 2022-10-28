package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/domain/usecases"
	"github.com/vodeacloud/hr-api/pkg/common"
	"github.com/vodeacloud/hr-api/pkg/pagination"
)

type Usecase struct {
	userRepo    repositories.UserRepository
	companyRepo repositories.CompanyRepository
}

func NewUsecase(
	userRepo repositories.UserRepository,
	companyRepo repositories.CompanyRepository,
) usecases.UserUsecase {
	return &Usecase{
		userRepo:    userRepo,
		companyRepo: companyRepo,
	}
}

func (u *Usecase) GetUsersByQuery(ctx context.Context, r *entities.GetUsersByQueryRequest) (*entities.UsersResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	q := &entities.UsersQuery{
		Search:     r.Search,
		Sort:       r.Sort,
		Pagination: pagination.NewPagination(r.Page, r.Limit),
	}

	users, err := u.userRepo.GetUsersByQuery(ctx, q)
	if err != nil {
		return nil, err
	}

	if err = u.appendUserRelations(ctx, users); err != nil {
		return nil, err
	}

	return &entities.UsersResponse{
		Users:      users,
		Pagination: q.Pagination,
	}, nil
}

func (u *Usecase) GetUserByID(ctx context.Context, r *entities.GetUserByIDRequest) (*entities.UserResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByID(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	if err = u.appendUserRelations(ctx, []*entities.User{user}); err != nil {
		return nil, err
	}

	return &entities.UserResponse{
		User: user,
	}, nil
}

func (u *Usecase) CreateUser(ctx context.Context, r *entities.CreateUserRequest) (*entities.UserResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	hashedPwd, err := common.HashAndSalt(r.Password)
	if err != nil {
		return nil, err
	}

	userType := entities.UserTypeUser
	if r.Type != nil {
		userType = *r.Type
	}

	user := &entities.User{
		Serial:    uuid.NewString(),
		Name:      r.Name,
		Email:     r.Email,
		Password:  hashedPwd,
		Phone:     r.Phone,
		AvatarURL: r.AvatarURL,
		Type:      userType,
	}

	user.UserDetail = &entities.UserDetail{
		UserSerial:    user.Serial,
		CompanySerial: r.UserDetail.CompanySerial,
	}

	if err = u.validateUniqueUser(ctx, user); err != nil {
		return nil, err
	}

	if err = u.validateUserDetail(ctx, user.UserDetail); err != nil {
		return nil, err
	}

	if err = u.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	if err = u.userRepo.UpsertUserDetail(ctx, user.UserDetail); err != nil {
		return nil, err
	}

	if err = u.appendUserRelations(ctx, []*entities.User{user}); err != nil {
		return nil, err
	}

	return &entities.UserResponse{
		User: user,
	}, nil
}

func (u *Usecase) UpdateUser(ctx context.Context, r *entities.UpdateUserRequest) (*entities.UserResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByID(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	user.Name = r.Name
	user.Email = r.Email
	user.Phone = r.Phone
	user.AvatarURL = r.AvatarURL

	user.UserDetail = &entities.UserDetail{
		UserSerial:    user.Serial,
		CompanySerial: r.UserDetail.CompanySerial,
	}

	if err = u.validateUniqueUser(ctx, user); err != nil {
		return nil, err
	}

	if err = u.validateUserDetail(ctx, user.UserDetail); err != nil {
		return nil, err
	}

	if err = u.userRepo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	if err = u.userRepo.UpsertUserDetail(ctx, user.UserDetail); err != nil {
		return nil, err
	}

	if err = u.appendUserRelations(ctx, []*entities.User{user}); err != nil {
		return nil, err
	}

	return &entities.UserResponse{
		User: user,
	}, nil
}

func (u *Usecase) DeleteUser(ctx context.Context, r *entities.DeleteUserRequest) (*entities.UserResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	user, err := u.userRepo.DeleteUser(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	return &entities.UserResponse{
		User: user,
	}, nil
}

func (u *Usecase) DeleteUserBulk(ctx context.Context, r *entities.DeleteUserBulkRequest) (*entities.UsersResponseWithoutPagination, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	users, err := u.userRepo.DeleteUserBulk(ctx, r.IDs)
	if err != nil {
		return nil, err
	}

	return &entities.UsersResponseWithoutPagination{
		Users: users,
	}, nil
}
