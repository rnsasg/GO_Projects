# Clean Architecture 

INSTANCE=$(docker inspect --format="{{ .NetworkSettings.Networks.IPAddress }}" cassandra)
echo $INSTANCE



export CASSANDRA_KEYSPACE="tagging"
export CASSANDRA_HOST="172.18.0.2"

## References 

1. https://cassandra.apache.org/_/quickstart.html 
