package oauth

import (
	"context"
	"errors"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"github.com/vodeacloud/hr-api/pkg/gormutil"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(
	db *gorm.DB,
) repositories.OAuthClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (r *ClientRepository) GetOAuthClientsByQuery(_ context.Context, q *entities.OAuthClientsQuery) (clients []*entities.OAuthClient, err error) {
	searchable := []string{"oauth_clients.name"}
	qb := r.db.Where(gormutil.SearchLikeRight(r.db, q.Search, searchable))

	q.Pagination.Total, err = gormutil.Count(qb, entities.OAuthClient{})
	if err != nil {
		return
	}
	q.Pagination.SetPagination()

	sortable := map[int32]string{1: "oauth_clients.id", 2: "oauth_clients.name"}
	if err = qb.
		Scopes(gormutil.Paginate(q.Pagination.Limit, q.Pagination.Page)).
		Scopes(gormutil.Sort(q.Sort, sortable)).
		Find(&clients).Error; err != nil {
		return
	}

	return
}

func (r *ClientRepository) GetOAuthClientByID(_ context.Context, id int64) (*entities.OAuthClient, error) {
	var client entities.OAuthClient
	if err := r.db.First(&client, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("client id not found")
		}
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) GetOAuthClientByName(_ context.Context, name string) (*entities.OAuthClient, error) {
	var client entities.OAuthClient
	if err := r.db.Where("oauth_clients.name = ?", name).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("client name not found")
		}
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) GetOAuthClientByClientID(_ context.Context, clientID string) (*entities.OAuthClient, error) {
	var client entities.OAuthClient
	if err := r.db.Where("oauth_clients.client_id = ?", clientID).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errutil.NewNotFoundError("client id not found")
		}
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepository) CreateOAuthClient(_ context.Context, client *entities.OAuthClient) error {
	return r.db.Create(&client).Error
}

func (r *ClientRepository) UpdateOAuthClient(_ context.Context, client *entities.OAuthClient) error {
	return r.db.Save(&client).Error
}

func (r *ClientRepository) DeleteOAuthClient(ctx context.Context, id int64) (*entities.OAuthClient, error) {
	client, err := r.GetOAuthClientByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = r.db.Delete(client).Error; err != nil {
		return nil, err
	}

	return client, nil
}

func (r *ClientRepository) DeleteOAuthClientBulk(_ context.Context, ids []int64) ([]*entities.OAuthClient, error) {
	//skip immediately when ids is empty
	if len(ids) == 0 {
		return nil, nil
	}

	var clients []*entities.OAuthClient
	if err := r.db.Find(&clients, ids).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&clients, ids).Error; err != nil {
		return nil, err
	}

	return clients, nil
}
