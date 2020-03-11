package uses_cases

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token/service"
)

type GetAccessTokenPort interface {
	Handler(id string) error
}
type GetAccessTokenUseCase struct {
	AccessTokenService service.Port
}

func (h *GetAccessTokenUseCase) Handler(id string) error {
	_, err := h.AccessTokenService.GetById(id)
	return err
}
