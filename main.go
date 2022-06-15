package main

import (
	"log"
	"supreme-go/account"
	"supreme-go/redis"
	"supreme-go/router"
	"supreme-go/settings"

	goredis "github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configuration := settings.Configuration{}
	redisClient := goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		Password: "P@ssw0rd",
		DB:       0,
	})
	redisStorage := redis.NewStorage(redisClient)
	accountService := account.NewService(redisStorage)
	httpRouter := router.NewHTTP(fiber.New(), accountService)

	if err := configuration.Load(); err != nil {
		log.Fatal(err)
	}

	if err := httpRouter.Route(configuration.Listen); err != nil {
		log.Fatal(err)
	}
}
