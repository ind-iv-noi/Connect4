package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var wg sync.WaitGroup
var db *sql.DB

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "indreias@LEAGUEINC"
	dbname   = "postgres"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, _ = sql.Open("postgres", psqlInfo)
	err := db.Ping()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	apiServer := httprouter.New()
	apiServer.GET("/api/start", startHandler)
	apiServer.POST("/api/move", moveHandler)
	apiServer.GET("/api/read", readHandler)
	apiServer.POST("/api/restart", restartHandler)
	apiServer.POST("/api/delete", deleteHandler)
	apiServer.GET("/api/turn", turnHandler)
	wg.Add(1)
	go func() {
		http.ListenAndServe(":8800", apiServer)
		wg.Done()
	}()
	wg.Wait()
}
