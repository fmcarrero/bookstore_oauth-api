CQL="CREATE  KEYSPACE oauth WITH REPLICATION = {'class' : 'SimpleStrategy' , 'replication_factor':1};
use oauth ;
CREATE TABLE access_tokens(access_token  varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);"

until echo $CQL | cqlsh; do
  echo "cqlsh: Cassandra is unavailable to initialize - will retry later"
  sleep 2
done &

exec /docker-entrypoint.sh "$@"