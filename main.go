package main

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/twn39/butterfly/handler"
	"github.com/twn39/butterfly/middleware"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

func getRedisClient(addr string, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return client
}

func getSettings() middleware.Settings {

	settings := middleware.Settings{}
	content, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &settings)
	if err != nil {
		panic(err)
	}

	return settings
}

func ServerStatic(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func main() {

	settings := getSettings()
	addr := settings.Redis.Host + ":" + settings.Redis.Port
	client := getRedisClient(addr, settings.Redis.Password, settings.Redis.Database)
	defer client.Close()

	route := mux.NewRouter()
	route.HandleFunc("/", handler.IndexHandler)
	route.HandleFunc("/posts/{id:[0-9]+}", handler.DetailHandler)
	route.Use(middleware.ConfigMiddleWare(&settings))
	route.Use(middleware.CacheMiddleWare(client))
	route.Use(middleware.GithubClientClientMiddleWare(settings.Github.Owner, settings.Github.Repo))
	ServerStatic(route)

	log.Fatal(http.ListenAndServe(":8080", route))
}
