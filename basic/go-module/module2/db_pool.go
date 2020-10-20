package main

import (
	"go-module/module2/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 5
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Println("Close: Connection", db.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func PerformQueries(query int, p *pool.Pool)  {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)

	// 模拟SQL查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Second)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			PerformQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("shutdown program")
	p.Close()
}