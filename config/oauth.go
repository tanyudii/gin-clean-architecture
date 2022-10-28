package config

import (
	"os"
)

type OAuthSecret struct {
	JWTPrivateKey string `envconfig:"JWT_PRIVATE_KEY" default:""`
	JWTPublicKey  string `envconfig:"JWT_PUBLIC_KEY" default:""`
}

const (
	pathStoragePrivateKey = "storage/keys/jwt-private.key"
	pathStoragePublicKey  = "storage/keys/jwt-public.key"
)

func (o *OAuthSecret) validateKey() {
	if o.JWTPrivateKey == "" {
		if file, _ := os.ReadFile(pathStoragePrivateKey); file != nil {
			o.JWTPrivateKey = string(file)
		}
	}
	if o.JWTPublicKey == "" {
		if file, _ := os.ReadFile(pathStoragePublicKey); file != nil {
			o.JWTPublicKey = string(file)
		}
	}
}
