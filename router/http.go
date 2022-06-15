package router

import (
	"supreme-go/account"

	"github.com/gofiber/fiber/v2"
)

type HTTP struct {
	router  *fiber.App
	account account.Service
}

func (r *HTTP) Route(listen string) error {
	r.router.Get("/account", r.account.HandleList)
	r.router.Get("/account/:id", r.account.HandleRead)
	r.router.Post("/account" /*r.account.ReadBody, r.account.Validate, r.account.WriteToRedis, */, r.account.HandleCreate)
	r.router.Patch("/account/:id", r.account.HandleUpdate)
	r.router.Delete("/account/:id", r.account.HandleRemove)

	return r.router.Listen(listen)
}

func NewHTTP(router *fiber.App, account account.Service) HTTP {
	return HTTP{
		router:  router,
		account: account,
	}
}
