package user

import (
	"context"
	"errors"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"github.com/vodeacloud/hr-api/pkg/gormutil"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) repositories.UserRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUsersByQuery(ctx context.Context, q *entities.UsersQuery) (users []*entities.User, err error) {
	searchable := []string{"users.name", "users.email"}
	qb := r.db.Where(gormutil.SearchLikeRight(r.db, q.Search, searchable))

	q.Pagination.Total, err = gormutil.Count(qb, entities.User{})
	if err != nil {
		return
	}
	q.Pagination.SetPagination()

	sortable := map[int32]string{1: "users.id", 2: "users.name", 3: "users.email"}
	if err = qb.
		Scopes(gormutil.Paginate(q.Pagination.Limit, q.Pagination.Page)).
		Scopes(gormutil.Sort(q.Sort, sortable)).
		Find(&users).Error; err != nil {
		return
	}

	return
}

func (r *Repository) GetUsersBySerials(_ context.Context, serials []string) ([]*entities.User, error) {
	//skip immediately when serials are empty
	if len(serials) == 0 {
		return nil, nil
	}
	var users []*entities.User
	if err := r.db.Where("users.serial IN (?)", serials).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetUserByID(_ context.Context, id int64) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("user id not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserBySerial(_ context.Context, serial string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("users.serial = ?", serial).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("user serial not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByName(_ context.Context, name string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("users.name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("user name not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByEmail(_ context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("users.email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("user email not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("users.phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("user phone not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(_ context.Context, user *entities.User) error {
	return r.db.Create(&user).Error
}

func (r *Repository) UpdateUser(_ context.Context, user *entities.User) error {
	return r.db.Save(&user).Error
}

func (r *Repository) DeleteUser(ctx context.Context, id int64) (*entities.User, error) {
	user, err := r.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = r.db.Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) DeleteUserBulk(_ context.Context, ids []int64) ([]*entities.User, error) {
	//skip immediately when ids are empty
	if len(ids) == 0 {
		return nil, nil
	}
	var users []*entities.User
	if err := r.db.Find(&users, ids).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(&users, ids).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetUserDetailsByUserSerials(ctx context.Context, userSerials []string) ([]*entities.UserDetail, error) {
	//skip immediately when userSerials are empty
	if len(userSerials) == 0 {
		return nil, nil
	}
	var userDetails []*entities.UserDetail
	if err := r.db.Where("user_details.user_serial IN (?)", userSerials).Find(&userDetails).Error; err != nil {
		return nil, err
	}
	return userDetails, nil
}

func (r *Repository) GetMapUserDetailByUserSerial(ctx context.Context, userSerials []string) (map[string]*entities.UserDetail, error) {
	userDetails, err := r.GetUserDetailsByUserSerials(ctx, userSerials)
	if err != nil {
		return nil, err
	}

	mapUserDetailByUserSerial := map[string]*entities.UserDetail{}
	for _, userDetail := range userDetails {
		mapUserDetailByUserSerial[userDetail.UserSerial] = userDetail
	}

	return mapUserDetailByUserSerial, nil
}

func (r *Repository) UpsertUserDetail(_ context.Context, userDetail *entities.UserDetail) error {
	return r.db.Save(userDetail).Error
}
