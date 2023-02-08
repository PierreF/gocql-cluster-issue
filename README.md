
Start the Cassandra:
```
docker compose up -d cassandra1 cassandra2 cassandra3
```

Wait for Cassandra to start. Use nodetool to wait for 3 node Up and Normal ("UN"):
```
docker compose exec cassandra1 nodetool status
```

Start the Go client:
```
docker compose up --build goclient 
```

Start the Python client:
```
docker compose up --build pyclient 
```

Everything works. Now start changing IP of Cassandra, one Cassandra by one. The following automate it:
```
for i in 1 2 3; do
    # Switch between 172.28.0.5N <-> 172.28.0.5N
    if grep -q "172.28.0.5$i" docker-compose.yml > /dev/null; then
        sed "s/172.28.0.5$i/172.28.0.6$i/" < docker-compose.yml > docker-compose.yml2
    else
        sed "s/172.28.0.6$i/172.28.0.5$i/" < docker-compose.yml > docker-compose.yml2
    fi
    mv  docker-compose.yml2  docker-compose.yml

    echo "Restarted Cassandra$i at $(date)"
    docker compose up -d cassandra$i
    
    # For for Cassandra to be UP. Ideally we should check using nodetool
    sleep 60
done
```

When the 3rd Cassandra is updated, Go client will fail and never recover with:
```
2023/02/08 15:32:17 gocql: no hosts available in the pool
```
