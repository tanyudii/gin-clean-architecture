package entities

import "github.com/vodeacloud/hr-api/pkg/errutil"

type OAuthCreateTokenResponse struct {
	AccessToken  string `form:"access_token"`
	RefreshToken string `form:"refresh_token,omitempty"`
	ExpiresIn    int32  `form:"expires_in"`
	TokenType    string `form:"token_type"`
	Scope        string `form:"scope"`
}

type OAuthCreateTokenRequest struct {
	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	GrantType    string `form:"grant_type"`
	Username     string `form:"username"`
	Password     string `form:"password"`
	Refresh      string `form:"refresh"`
	Scope        string `form:"scope"`
}

func (r *OAuthCreateTokenRequest) Validate() error {
	fields := errutil.ErrorField{}
	if r.ClientID == "" {
		fields["client_id"] = "client field is required"
	}
	if r.ClientSecret == "" {
		fields["client_id"] = "client secret field is required"
	}
	if r.GrantType == "" {
		fields["grant_type"] = "grant type field is required"
	}
	switch r.GrantType {
	case "password":
		r.validateGrantTypePassword(fields)
	case "refresh_token":
		r.validateGrantTypeRefreshToken(fields)
	default:
		return ErrUnsupportedGrantType
	}
	return nil
}

func (r *OAuthCreateTokenRequest) validateGrantTypePassword(fields errutil.ErrorField) {
	if r.Username == "" {
		fields["username"] = "username field is required"
	}
	if r.Password == "" {
		fields["password"] = "password field is required"
	}
}

func (r *OAuthCreateTokenRequest) validateGrantTypeRefreshToken(fields errutil.ErrorField) {
	if r.Refresh == "" {
		fields["refresh"] = "refresh field is required"
	}
}
