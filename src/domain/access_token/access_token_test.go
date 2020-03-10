package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessTokenConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24")
}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.NotNil(t, at)
	assert.EqualValues(t, "", at.AccessToken, "accessToken not should be empty")
	assert.EqualValues(t, 0, at.UserId, "userId not should be empty")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired())
}
