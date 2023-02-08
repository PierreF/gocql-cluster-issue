package main

import (
	"context"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("cassandra1:9042", "cassandra2:9042")
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	ctx := context.Background()

	for ctx.Err() == nil {
		var version string
		if err := session.Query("SELECT release_version FROM system.local").WithContext(ctx).Scan(&version); err != nil {
			log.Print(err)
		}
		log.Println("release_version=", version)
		time.Sleep(15 * time.Second)
	}
}
