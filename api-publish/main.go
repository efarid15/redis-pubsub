package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notification/config"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

type Publish struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

func publishMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := Publish{
		Channel: r.FormValue("channel"),
		Message: r.FormValue("message"),
	}

	opt, err := redis.ParseURL(os.Getenv("REDIS_SERVER"))
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	rdb.Publish(ctx, r.FormValue("channel"), r.FormValue("message")).Err()

	json.NewEncoder(w).Encode(payload)

}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/publish", publishMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

func main() {
	config.LoadFile(".env")
	fmt.Println("Rest API v2.0 - Redis Pub/Sub")
	handleRequest()
}
