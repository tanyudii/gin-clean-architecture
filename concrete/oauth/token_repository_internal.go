package oauth

import "fmt"

const (
	prefixAccessToken  = "AccessToken"
	prefixRefreshToken = "RefreshToken"
)

func keyAccessToken(id string) string {
	return fmt.Sprintf("%s:%s", prefixAccessToken, id)
}

func keyRefreshToken(id string) string {
	return fmt.Sprintf("%s:%s", prefixRefreshToken, id)
}
