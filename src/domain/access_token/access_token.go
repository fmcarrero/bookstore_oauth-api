package access_token

import (
	"errors"
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/domain/utils/crypto_utils"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() error {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.New("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.New("invalid users id")
	}
	if at.ClientId <= 0 {
		return errors.New("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.New("invalid expiration time")
	}
	return nil
}
func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
