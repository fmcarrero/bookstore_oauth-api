package repository

import (
	"errors"
	"fmt"
	"github.com/fmcarrero/bookstore_oauth-api/src/domain/access_token"
	"github.com/fmcarrero/bookstore_oauth-api/src/infrastructure/cassandra"
	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires  FROM access_tokens WHERE Access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES ( ?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires = ? where access_token = ?;"
)

type AccessTokenCassandraRepository struct {
}

func (repository *AccessTokenCassandraRepository) GetById(id string) (*access_token.AccessToken, error) {

	var accessToken access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&accessToken.AccessToken, &accessToken.UserId, &accessToken.ClientId, &accessToken.Expires); err != nil {
		fmt.Println(err)
		if err == gocql.ErrNotFound {
			return nil, errors.New(fmt.Sprintf("no access token when give id %s", id))
		}
		return nil, errors.New(fmt.Sprintf("error when access data %s", id))
	}

	return &accessToken, nil
}

func (repository *AccessTokenCassandraRepository) Create(token *access_token.AccessToken) error {

	if err := cassandra.GetSession().Query(queryCreateAccessToken, token.AccessToken, token.UserId, token.ClientId, token.Expires).Exec(); err != nil {
		logger.Error(err.Error(), err)
		return errors.New(fmt.Sprint("error when created access token"))
	}
	return nil
}

func (repository *AccessTokenCassandraRepository) UpdateExpirationTime(token access_token.AccessToken) error {

	if err := cassandra.GetSession().Query(queryUpdateExpires, token.Expires, token.AccessToken).Exec(); err != nil {
		fmt.Println(err)
		return errors.New(fmt.Sprint("error when updated access token"))
	}
	return nil
}
