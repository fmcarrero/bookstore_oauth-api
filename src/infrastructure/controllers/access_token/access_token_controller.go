package access_token

import (
	"fmt"
	"github.com/fmcarrero/bookstore_oauth-api/src/application/uses_cases"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token_request"
	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/fmcarrero/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RedirectAccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	UpdateTimeExpires(c *gin.Context)
}
type Handler struct {
	GetAccessTokenUseCase    uses_cases.GetAccessTokenPort
	CreateAccessTokenUseCase uses_cases.CreateAccessTokenPort
	UpdateTimeExpiresUseCase uses_cases.UpdateSetExpiresPort
}

func (h *Handler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.GetAccessTokenUseCase.Handler(accessTokenId)
	if err != nil {
		if strings.Contains(err.Error(), "no access token when give") {
			restErr := rest_errors.NewNotFoundError(fmt.Sprintf("%s not found", accessTokenId))
			logger.Error(restErr.Message(), restErr)
			c.JSON(restErr.Status(), restErr.Message())
			return
		}
		restErr := rest_errors.NewBadRequestError(fmt.Sprintf("bad request with id %s", accessTokenId))
		logger.Error(restErr.Message(), restErr)
		c.JSON(restErr.Status(), restErr.Message())
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *Handler) Create(c *gin.Context) {
	var request access_token_request.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	token, err := h.CreateAccessTokenUseCase.Handler(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, token)
}
func (h *Handler) UpdateTimeExpires(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.GetAccessTokenUseCase.Handler(accessTokenId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
