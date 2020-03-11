package uses_cases

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token/service"
)

type UpdateSetExpiresPort interface {
	Handler(token access_token.AccessToken) error
}
type UpdateSetExpiresUseCase struct {
	AccessTokenService service.AccessTokenServicePort
}

func (h *UpdateSetExpiresUseCase) Handler(token access_token.AccessToken) error {
	return h.AccessTokenService.UpdateExpirationTime(token)

}
