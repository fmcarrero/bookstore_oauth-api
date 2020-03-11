package users

import (
	"errors"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/users"
	"github.com/go-resty/resty/v2"
)

var (
	UserRestClient = resty.New()
)

type UserRestRepository struct {
}

func (r *UserRestRepository) LoginUser(email string, password string) (*users.User, error) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	var user users.User
	var errLogin error
	_, errPost := UserRestClient.SetHostURL("http://localhost:8080").R().
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
