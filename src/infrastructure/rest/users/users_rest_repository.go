package users

import (
	"errors"
	"fmt"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/users"
	"github.com/go-resty/resty/v2"
	"os"
)

var (
	UserRestClient = resty.New()
)

type UserRestRepository struct {
}

func (r *UserRestRepository) LoginUser(email string, password string) (*users.User, error) {
	hostUserServiceHost := os.Getenv("USERS_API_HOST")
	hostUserServicePort := fmt.Sprintf(":%s", os.Getenv("USERS_API_PORT"))
	host := fmt.Sprintf("http://%s%s", hostUserServiceHost, hostUserServicePort)
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	var user users.User
	var errLogin error
	_, errPost := UserRestClient.SetHostURL(host).R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetError(&errLogin).
		SetResult(&user).
		Post("/users/login")
	if errPost != nil {
		return nil, errors.New(errPost.Error())
	}
	if errLogin != nil {
		return nil, errors.New(errLogin.Error())
	}
	return &user, nil
}
