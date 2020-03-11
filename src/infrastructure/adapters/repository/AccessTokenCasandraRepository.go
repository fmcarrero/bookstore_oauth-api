package repository

import (
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/adapters/repository/cassandra"
)

type AccessTokenCassandraRepository struct {
}

func (acc *AccessTokenCassandraRepository) GetById(string) (*access_token.AccessToken, error) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, nil
}
