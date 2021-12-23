package router

import (
	"screenshot-site/handle"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	Fiber  *fiber.App
	Handle handle.ScreenShortHandle
}

func (api *API) SetupRouter() {
	app := api.Fiber.Group("/")
	app.Get("fullpage", api.Handle.ScreenShot)
}
