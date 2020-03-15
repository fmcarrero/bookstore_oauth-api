package cassandra

import (
	"github.com/gocql/gocql"
	"os"
	"strconv"
)

var (
	session *gocql.Session
	hostEnv = "CASSANDRA_HOST"
	portEnv = "CASSANDRA_PORT"
)

func Run() {
	// connect to Cassandra the cluster
	port, _ := strconv.ParseInt(os.Getenv(portEnv), 10, 64)

	cluster := gocql.NewCluster(os.Getenv(hostEnv))
	cluster.Port = int(port)
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

}

func GetSession() *gocql.Session {
	return session
}
