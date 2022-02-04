package main

import (
	"fmt"

	"github.com/config"
	"github.com/gofiber/fiber/v2"
	"github.com/models"
	"github.com/routes"
	"github.com/spf13/viper"
)

func init() {
	config.SetupVars()
	models.SetupDB()
}
func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(fmt.Sprintf(":%d", viper.GetInt("port")))
}
