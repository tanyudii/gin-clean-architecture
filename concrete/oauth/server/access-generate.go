package oauth

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/golang-jwt/jwt"
	"github.com/vodeacloud/hr-api/config"
	"github.com/vodeacloud/hr-api/domain/entities"
)

type InternalAccessGenerate interface {
	oauth2.AccessGenerate
	VerifyAccess(ctx context.Context, access string) (*entities.AccessTokenClaims, error)
	VerifyRefresh(ctx context.Context, refresh string) (*entities.RefreshTokenClaims, error)
}

type accessGenerate struct {
	cfg *config.Config
}

func newAccessGenerate(cfg *config.Config) InternalAccessGenerate {
	return &accessGenerate{
		cfg: cfg,
	}
}

func (g *accessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
	jwtPrivateKey, err := g.getPrivateKey()
	if err != nil {
		return "", "", err
	}

	access, err = g.createAccessToken(ctx, data, jwtPrivateKey)
	if err != nil {
		return
	}

	if isGenRefresh {
		refresh, err = g.createRefreshToken(ctx, data, jwtPrivateKey)
		if err != nil {
			return
		}
	}

	return
}

func (g *accessGenerate) VerifyAccess(_ context.Context, access string) (*entities.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(access, &entities.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, entities.ErrInvalidOrExpiredAccess
		}
		return g.getPublicKey()
	})
	if err != nil {
		return nil, entities.ErrInvalidOrExpiredAccess
	}

	claims, ok := token.Claims.(*entities.AccessTokenClaims)
	if !ok {
		return nil, entities.ErrInvalidOrExpiredAccess
	}

	return claims, nil
}

func (g *accessGenerate) VerifyRefresh(_ context.Context, refresh string) (*entities.RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(refresh, &entities.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, entities.ErrInvalidOrExpiredRefresh
		}
		return g.getPublicKey()
	})
	if err != nil {
		return nil, entities.ErrInvalidOrExpiredRefresh
	}

	claims, ok := token.Claims.(*entities.RefreshTokenClaims)
	if !ok {
		return nil, entities.ErrInvalidOrExpiredRefresh
	}

	return claims, nil
}
