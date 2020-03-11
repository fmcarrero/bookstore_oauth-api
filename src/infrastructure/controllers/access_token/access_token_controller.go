package access_token

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/application/uses_cases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RedirectAccessTokenHandler interface {
	GetById(c *gin.Context)
}
type Handler struct {
	GetAccessTokenUseCase uses_cases.GetAccessTokenPort
}

func (h *Handler) GetById(c *gin.Context) {
	accessToken := strings.TrimSpace(c.Param("access_token_id"))
	err := h.GetAccessTokenUseCase.Handler(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, "implement me")
	}
	c.JSON(http.StatusNotImplemented, "implement me")
}
