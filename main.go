package main

import (
	"os"

	// TODO: Buat reponya biar bisa diinstall
	"github.com/nadiastore/go-api/pkg/configs"
	"github.com/nadiastore/go-api/pkg/middleware"
	"github.com/nadiastore/go-api/pkg/routes"
	"github.com/nadiastore/go-api/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	if os.Getenv("SERVER_ENV") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
