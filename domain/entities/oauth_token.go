package entities

import (
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/golang-jwt/jwt"
	"time"
)

type AccessToken struct {
	ID             string             `json:"id"`
	ClientID       string             `json:"clientID"`
	IssuedAt       time.Time          `json:"issuedAt"`
	ExpiresAt      time.Time          `json:"expiresAt"`
	StandardClaims jwt.StandardClaims `json:"standardClaims"`
	TokenInfo      *models.Token      `json:"tokenInfo,omitempty"`
}

type AccessTokenClaims struct {
	jwt.StandardClaims
}

type RefreshToken struct {
	ID             string             `json:"id"`
	ClientID       string             `json:"clientID"`
	IssuedAt       time.Time          `json:"issuedAt"`
	ExpiresAt      time.Time          `json:"expiresAt"`
	StandardClaims jwt.StandardClaims `json:"standardClaims"`
	TokenInfo      *models.Token      `json:"tokenInfo,omitempty"`
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
}
