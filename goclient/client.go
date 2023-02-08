package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gocql/gocql"
)

var (
	mutex         sync.Mutex
	reopenSession bool
)

type observer struct{}

func (observer) ObserveConnect(info gocql.ObservedConnect) {
	if info.Err != nil {
		mutex.Lock()
		reopenSession = true
		mutex.Unlock()
	}
}

func main() {
	cluster := gocql.NewCluster("cassandra1:9042", "cassandra2:9042")
	cluster.ConnectObserver = observer{}
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	ctx := context.Background()

	for ctx.Err() == nil {
		// We can't hold the lock during session re-openning. gocql will call ObserveConnect and dead-lock if we hold it.
		mutex.Lock()
		needReOpen := reopenSession
		mutex.Unlock()

		if needReOpen {
			log.Println("reopen session")
			session.Close()
			session, err = cluster.CreateSession()
			if err != nil {
				log.Fatal(err)
			}

			mutex.Lock()
			reopenSession = false
			mutex.Unlock()
		}

		var version string
		if err := session.Query("SELECT release_version FROM system.local").WithContext(ctx).Scan(&version); err != nil {
			log.Print(err)
		}
		log.Println("release_version=", version)
		time.Sleep(15 * time.Second)
	}
}
