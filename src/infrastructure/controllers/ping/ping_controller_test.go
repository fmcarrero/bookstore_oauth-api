// +build integration

package ping_test

import (
	"context"
	"fmt"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/app"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/controllers/ping"
	"github.com/fmcarrero/bookstore_oauth-api/src/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	r = app.Router
)

func TestMain(m *testing.M) {
	fmt.Println("about to start oauth tests")
	containerMockServer, ctx := utils.LoadCassandra()
	code := m.Run()
	beforeAll(containerMockServer, ctx)
	os.Exit(code)
}

func TestPing(t *testing.T) {

	r.GET("/ping", ping.Ping)

	w := performRequest(r, "GET", "/ping", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())

}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func beforeAll(container testcontainers.Container, ctx context.Context) {
	_ = container.Terminate(ctx)
}
