package oauth

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/vodeacloud/hr-api/domain/entities"
	"github.com/vodeacloud/hr-api/domain/repositories"
	"github.com/vodeacloud/hr-api/pkg/errutil"
)

type tokenStore struct {
	tokenRepo repositories.OAuthTokenRepository
	accessGen InternalAccessGenerate
}

func newTokenStore(
	tokenRepo repositories.OAuthTokenRepository,
	accessGen InternalAccessGenerate,
) oauth2.TokenStore {
	return &tokenStore{
		tokenRepo: tokenRepo,
		accessGen: accessGen,
	}
}

func (s *tokenStore) Create(ctx context.Context, info oauth2.TokenInfo) error {
	atClaims, err := s.accessGen.VerifyAccess(ctx, info.GetAccess())
	if err != nil {
		return err
	}

	accessToken := &entities.AccessToken{
		ID:             atClaims.StandardClaims.Id,
		ClientID:       info.GetClientID(),
		IssuedAt:       info.GetAccessCreateAt(),
		ExpiresAt:      info.GetAccessCreateAt().Add(info.GetAccessExpiresIn()),
		StandardClaims: atClaims.StandardClaims,
		TokenInfo:      info.(*models.Token),
	}
	if err = s.tokenRepo.CreateAccessToken(ctx, accessToken); err != nil {
		return err
	}

	if refresh := info.GetRefresh(); refresh != "" {
		refreshClaims, err := s.accessGen.VerifyRefresh(ctx, refresh)
		if err != nil {
			return err
		}

		refreshToken := &entities.RefreshToken{
			ID:             refreshClaims.StandardClaims.Id,
			ClientID:       info.GetClientID(),
			IssuedAt:       info.GetRefreshCreateAt(),
			ExpiresAt:      info.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn()),
			StandardClaims: atClaims.StandardClaims,
			TokenInfo:      info.(*models.Token),
		}
		if err = s.tokenRepo.CreateRefreshToken(ctx, refreshToken); err != nil {
			return err
		}
	}
	return nil
}

func (s *tokenStore) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	accessClaims, err := s.accessGen.VerifyAccess(ctx, access)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokenRepo.GetAccessToken(ctx, accessClaims.Id)
	if err != nil {
		if errutil.IsNotFoundError(err) {
			err = entities.ErrInvalidOrExpiredAccess
		}
		return nil, err
	}

	return accessToken.TokenInfo, nil
}

func (s *tokenStore) RemoveByAccess(ctx context.Context, access string) error {
	accessClaims, err := s.accessGen.VerifyAccess(ctx, access)
	if err != nil {
		return err
	}

	return s.tokenRepo.DeleteAccessToken(ctx, accessClaims.Id)
}

func (s *tokenStore) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	refreshClaims, err := s.accessGen.VerifyRefresh(ctx, refresh)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenRepo.GetRefreshToken(ctx, refreshClaims.Id)
	if err != nil {
		if errutil.IsNotFoundError(err) {
			err = entities.ErrInvalidOrExpiredRefresh
		}
		return nil, err
	}

	return refreshToken.TokenInfo, nil
}

func (s *tokenStore) RemoveByRefresh(ctx context.Context, refresh string) error {
	refreshClaims, err := s.accessGen.VerifyRefresh(ctx, refresh)
	if err != nil {
		return err
	}

	return s.tokenRepo.DeleteRefreshToken(ctx, refreshClaims.Id)
}

func (s *tokenStore) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	panic("implement me")
}

func (s *tokenStore) RemoveByCode(ctx context.Context, code string) error {
	panic("implement me")
}
