// +build integration

package access_token_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	at "github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token_request"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/adapters/repository"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/app"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/cassandra"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/rest/users"
	"github.com/fmcarrero/bookstore_oauth-api/src/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	r               = app.Router
	handler         = access_token.Handler{}
	pathCreateToken = "/oauth/access_token"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start oauth tests")
	loadHandler()

	containerCassandra, ctxCassandra := utils.LoadCassandra()
	cassandra.Run()
	containerUserHost, ctxUser := utils.LoadUserService()
	code := m.Run()
	beforeAllCassandra(containerCassandra, ctxCassandra)
	beforeAllUsers(containerUserHost, ctxUser)
	os.Exit(code)
}
func loadHandler() {
	service := app.CreateAccessTokenService(&repository.AccessTokenCassandraRepository{}, &users.UserRestRepository{})
	createAccessToken := app.CreateAccessTokenUseCase(service)
	handler = access_token.Handler{
		CreateAccessTokenUseCase: createAccessToken,
	}

	r.POST(pathCreateToken, handler.Create)
}

func TestHandler_Create(t *testing.T) {
	accessToken := access_token_request.AccessTokenRequest{
		GrantType:    "password",
		Username:     "mauriciocarrero15@gmail.com",
		Password:     "sistemas31",
		ClientId:     "1",
		ClientSecret: "1",
	}
	body, _ := json.Marshal(accessToken)

	w := performRequest(r, "POST", pathCreateToken, bytes.NewReader(body))

	var response at.AccessToken
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.EqualValues(t, 1, response.UserId)
	assert.NotEmpty(t, response.AccessToken, "access should has a value")
	assert.NotEmpty(t, response.Expires, "expires should has a value")
}

func TestHandler_Create_Invalid_Body(t *testing.T) {

	w := performRequest(r, "POST", pathCreateToken, ioutil.NopCloser(bytes.NewReader([]byte(""))))

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid json body")

}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func beforeAllCassandra(container testcontainers.Container, ctx context.Context) {
	_ = container.Terminate(ctx)
}
func beforeAllUsers(container testcontainers.Container, ctx context.Context) {
	_ = container.Terminate(ctx)
}
