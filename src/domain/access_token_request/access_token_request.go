package access_token_request

import "errors"

const (
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (attReq *AccessTokenRequest) Validate() error {
	switch attReq.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
		break
	default:
		return errors.New("invalid grant type parameter")
	}
	return nil
}
