package ports

import users "github.com/fmcarrero/bookstore_oauth-api/src/domain/users"

type UserRepository interface {
	LoginUser(email string, password string) (*users.User, error)
}
