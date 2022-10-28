package role

import (
	"context"
	"errors"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"github.com/vodeacloud/hr-api/pkg/gormutil"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) repositories.RoleRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetRolesByQuery(_ context.Context, q *entities.RolesQuery) (roles []*entities.Role, err error) {
	searchable := []string{"roles.name"}
	qb := r.db.Where(gormutil.SearchLikeRight(r.db, q.Search, searchable))

	q.Pagination.Total, err = gormutil.Count(qb, entities.Role{})
	if err != nil {
		return
	}
	q.Pagination.SetPagination()

	sortable := map[int32]string{1: "roles.id", 2: "roles.name"}
	if err = qb.
		Scopes(gormutil.Paginate(q.Pagination.Limit, q.Pagination.Page)).
		Scopes(gormutil.Sort(q.Sort, sortable)).
		Find(&roles).Error; err != nil {
		return
	}

	return
}

func (r *Repository) GetRolesBySerials(_ context.Context, serials []string) ([]*entities.Role, error) {
	//skip immediately when serials are empty
	if len(serials) == 0 {
		return nil, nil
	}
	var roles []*entities.Role
	if err := r.db.Where("roles.serial IN (?)", serials).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *Repository) GetRoleByID(_ context.Context, id int64) (*entities.Role, error) {
	var role entities.Role
	if err := r.db.First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("role id not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *Repository) GetRoleBySerial(_ context.Context, serial string) (*entities.Role, error) {
	var role entities.Role
	if err := r.db.Where("roles.serial = ?", serial).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("role serial not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *Repository) GetRoleByName(_ context.Context, name string) (*entities.Role, error) {
	var role entities.Role
	if err := r.db.Where("roles.name = ?", name).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("role name not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *Repository) CreateRole(_ context.Context, role *entities.Role) error {
	return r.db.Create(&role).Error
}

func (r *Repository) UpdateRole(_ context.Context, role *entities.Role) error {
	return r.db.Save(&role).Error
}

func (r *Repository) DeleteRole(ctx context.Context, id int64) (*entities.Role, error) {
	role, err := r.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = r.db.Delete(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *Repository) DeleteRoleBulk(_ context.Context, ids []int64) ([]*entities.Role, error) {
	//skip immediately when ids are empty
	if len(ids) == 0 {
		return nil, nil
	}
	var roles []*entities.Role
	if err := r.db.Find(&roles, ids).Error; err != nil {
		return nil, err
	}
	if err := r.db.Delete(&roles, ids).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *Repository) GetRolePermissionsByRoleSerials(_ context.Context, roleSerials []string) ([]*entities.RolePermission, error) {
	var rolePermissions []*entities.RolePermission
	if len(roleSerials) == 0 {
		return rolePermissions, nil
	}

	if err := r.db.Where("role_permissions.role_serial IN (?)", roleSerials).Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (r *Repository) SyncRolePermission(_ context.Context, role *entities.Role, rolePermissions []*entities.RolePermission) error {
	var permissionCodes []string
	for _, userRole := range rolePermissions {
		permissionCodes = append(permissionCodes, userRole.PermissionCode)
		userRole.RoleSerial = role.Serial
	}

	qRemoveRolePermissions := r.db.Where("role_permissions.role_serial = ?", role.Serial)
	if len(permissionCodes) != 0 {
		qRemoveRolePermissions.Where("role_permissions.permission_code NOT IN (?)", permissionCodes)
	}

	if err := qRemoveRolePermissions.Delete(&entities.RolePermission{}).Error; err != nil {
		return err
	}

	// skip immediately if user roles empty
	if len(rolePermissions) == 0 {
		return nil
	}

	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&rolePermissions).Error
}
