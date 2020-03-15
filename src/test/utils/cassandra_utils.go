package utils

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
)

func LoadCassandra() (testcontainers.Container, context.Context) {

	schemaImportCassandra, errReadFile := ioutil.ReadFile("../../../test/resources/import_cassandra.txt")
	if errReadFile != nil {
		panic(errReadFile)
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "cassandra:latest",
		ExposedPorts: []string{"9042/tcp", "9160/tcp"},
		WaitingFor:   wait.ForLog(" Starting listening for CQL clients on /0.0.0.0:9042 (unencrypted)..."),
	}

	cassandraC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		logger.Error(err.Error(), err)
		panic(err)
	}

	host, _ := cassandraC.Host(ctx)
	p, _ := cassandraC.MappedPort(ctx, "9042/tcp")

	_ = os.Setenv("CASSANDRA_HOST", host)
	_ = os.Setenv("CASSANDRA_PORT", p.Port())
	retURL := fmt.Sprintf("localhost:%s", p.Port())

	port, _ := strconv.Atoi(p.Port())
	clusterConfig := gocql.NewCluster(retURL)
	clusterConfig.ProtoVersion = 4
	clusterConfig.Port = port
	log.Printf("%v", clusterConfig.Port)

	session, err := clusterConfig.CreateSession()

	if err != nil {
		panic(fmt.Errorf("error creating session: %s", err))
	}
	errInsert := session.Query(string(schemaImportCassandra)).Exec()
	if errInsert != nil {
		panic(errInsert)
	}

	if err := session.Query("CREATE TABLE oauth.access_tokens(access_token  varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);").Exec(); err != nil {
		panic(err)
	}

	defer session.Close()

	return cassandraC, ctx
}
