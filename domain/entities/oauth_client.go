package entities

import (
	"github.com/vodeacloud/hr-api/pkg/common"
	"time"
)

type OAuthClient struct {
	ID           int64
	Name         string
	IsActive     bool
	IsInternal   bool
	ClientID     string
	ClientSecret string
	Domain       *string
	UserID       *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	CreatedBy    string
	UpdatedBy    string
}

func (o OAuthClient) TableName() string {
	return "oauth_clients"
}

func (o OAuthClient) GetID() string {
	return o.ClientID
}

func (o OAuthClient) GetSecret() string {
	return o.ClientSecret
}

func (o OAuthClient) GetDomain() string {
	if o.Domain == nil {
		return ""
	}
	return *o.Domain
}

func (o OAuthClient) GetUserID() string {
	if o.UserID == nil {
		return ""
	}
	return *o.UserID
}

func (o OAuthClient) VerifyPassword(val string) bool {
	valid, _ := common.CompareHashed(o.ClientSecret, val)
	return valid
}
