package oauth

import (
	"context"
	"crypto/rsa"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/vodeacloud/hr-api/domain/entities"
	"time"
)

func (g *accessGenerate) createAccessToken(_ context.Context, data *oauth2.GenerateBasic, privateKey *rsa.PrivateKey) (string, error) {
	data.TokenInfo.SetAccessExpiresIn(time.Hour * 24 * 7)
	atClaims := entities.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.NewString(),
			Subject:   data.UserID,
			IssuedAt:  data.TokenInfo.GetAccessCreateAt().Unix(),
			NotBefore: data.TokenInfo.GetAccessCreateAt().Unix(),
			ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	signedAt, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedAt, nil
}

func (g *accessGenerate) createRefreshToken(_ context.Context, data *oauth2.GenerateBasic, privateKey *rsa.PrivateKey) (string, error) {
	data.TokenInfo.SetRefreshExpiresIn(time.Hour * 24 * 90)
	rtClaims := jwt.StandardClaims{
		Id:        uuid.NewString(),
		Subject:   data.UserID,
		IssuedAt:  data.TokenInfo.GetRefreshCreateAt().Unix(),
		NotBefore: data.TokenInfo.GetRefreshCreateAt().Unix(),
		ExpiresAt: data.TokenInfo.GetRefreshCreateAt().Add(data.TokenInfo.GetRefreshExpiresIn()).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, rtClaims)
	signedRt, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedRt, nil
}

func (g *accessGenerate) getPrivateKey() (*rsa.PrivateKey, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(g.cfg.OAuthSecret.JWTPrivateKey))
	if err != nil {
		return nil, entities.ErrPrivateOrPublicKeyInvalid
	}
	return key, nil
}

func (g *accessGenerate) getPublicKey() (*rsa.PublicKey, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(g.cfg.OAuthSecret.JWTPublicKey))
	if err != nil {
		return nil, entities.ErrPrivateOrPublicKeyInvalid
	}
	return key, nil
}
