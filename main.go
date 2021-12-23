package main

import (
	"os"
	"site-screenshot/handle"
	"site-screenshot/router"

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
	screenShort := handle.ScreenShortHandle{}
	api := router.API{
		Fiber:  f,
		Handle: screenShort,
	}
	api.SetupRouter()
	f.Listen(":" + os.Getenv("PORT"))
}
