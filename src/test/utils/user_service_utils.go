package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	client = http.Client{Timeout: 5 * time.Second}
)

func LoadUserService() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mockserver/mockserver",
		ExposedPorts: []string{"1080/tcp"},
		WaitingFor:   wait.ForListeningPort("1080"),
	}
	containerMockServer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	ip, err := containerMockServer.Host(ctx)
	if err != nil {
		panic(err)
	}

	port, err := containerMockServer.MappedPort(ctx, "1080/tcp")
	if err != nil {
		panic(err)
	}
	_ = os.Setenv("PORT_OAUTH_SERVICE", port.Port())
	host := fmt.Sprintf("http://%s:%s", ip, port.Port())
	data, errReadFile := ioutil.ReadFile("../../../test/resources/request_validate_user_franklin_ok.json")
	if errReadFile != nil {
		panic(errReadFile)
	}
	requestBody := ioutil.NopCloser(bytes.NewReader(data))
	requestLoadInformation, errPut := http.NewRequest("PUT", host+"/expectation", requestBody)
	if errPut != nil {
		panic(errPut)
	}
	_, _ = client.Do(requestLoadInformation)
	_ = os.Setenv("USERS_API_HOST", ip)
	_ = os.Setenv("USERS_API_PORT", port.Port())
	return containerMockServer, ctx
}
