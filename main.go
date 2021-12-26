package main

import (
	"os"

	"screenshot-site/handle"
	"screenshot-site/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		return
	}
}

func main() {
	f := fiber.New()
	screenShort := handle.ScreenShotHandle{}
	api := router.API{
		Fiber:  f,
		Handle: screenShort,
	}
	api.SetupRouter()
	f.Listen(":" + os.Getenv("PORT"))
}
