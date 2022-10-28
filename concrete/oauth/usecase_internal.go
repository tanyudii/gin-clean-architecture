package oauth

import (
	"context"
	"errors"
	"github.com/go-oauth2/oauth2/v4"
	oautherr "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"strconv"
)

func (u *Usecase) grantTypePassword(ctx context.Context, r *entities.OAuthCreateTokenRequest, tokenGenReq *oauth2.TokenGenerateRequest) error {
	user, err := u.userUc.GetUserByEmailAndPassword(ctx, &entities.GetUserByEmailAndPasswordRequest{
		Email:    r.Username,
		Password: r.Password,
	})
	if err != nil {
		return err
	}
	tokenGenReq.UserID = strconv.FormatInt(user.ID, 10)
	return nil
}

func (u *Usecase) transformErrOAuth2(err error) error {
	if errors.Is(err, oautherr.ErrInvalidClient) {
		err = errutil.NewBadRequestError(err.Error())
	} else if errors.Is(err, oautherr.ErrInvalidAccessToken) {
		err = errutil.NewBadRequestError(err.Error())
	}
	return err
}

func (u *Usecase) validateClient(ctx context.Context, client *entities.OAuthClient) error {
	oauthClientByClientID, err := u.clientRepo.GetOAuthClientByClientID(ctx, client.ClientID)
	if err != nil {
		if errutil.IsNotFoundError(err) {
			err = entities.ErrInvalidClient
		}
		return err
	}

	if !oauthClientByClientID.VerifyPassword(client.ClientSecret) {
		return entities.ErrInvalidClient
	}

	return nil
}

func (u *Usecase) getRolesBySerials(ctx context.Context, serials []string) ([]*entities.Role, error) {
	if len(serials) == 0 {
		return nil, nil
	}

	roles, err := u.roleRepo.GetRolesBySerials(ctx, serials)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
