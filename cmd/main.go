package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/toshkentov01/template/api/routes"
	"github.com/toshkentov01/template/config"
	"github.com/toshkentov01/template/pkg/logger"
	"github.com/toshkentov01/template/pkg/middleware"
	"github.com/toshkentov01/template/pkg/migration"
	"github.com/toshkentov01/template/pkg/utils"

	_ "github.com/joho/godotenv/autoload" // autoload env
	_ "github.com/toshkentov01/template/api/docs" //register swagger

	service "github.com/toshkentov01/template/service/user"
)

// @title Monolith-Template 2.0
// @version 0.1
// @description This is an auto-generated API Docs for Monolith-Template 2.0.
// @termsOfService http://swagger.io/terms/
// @contact.name Sardor Toshkentov
// @contact.email toshkentovsardor.2003@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	// ============================================================ //
	// Migrations SetUp ...

	migration.Up()

	// ============================================================ //
	// Config Setup ...

	cfg := config.Config()
	fiberConfig := config.FiberConfig()

	// ============================================================ //
	// Fiber Setup ...

	app := fiber.New(fiberConfig)

	app.Use(fiberLogger.New(fiberLogger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	middleware.FiberMiddleware(app)

	// ============================================================ //
	// JWT Setup ...

	jwtRoleAuthorizer, err := middleware.NewJWTRoleAuthorizer(cfg)
	if err != nil {
		log.Fatal("Could not initialize JWT Role Authorizer")
	}

	app.Use(middleware.NewAuthorizer(jwtRoleAuthorizer))

	// ============================================================ //
	// Service SetUp ...

	l := logger.New(cfg.LogLevel, "monoliith")
	userService := service.NewUserService(l)

	// ============================================================ //
	// Api Setup ...

	routes.SwaggerRoute(app)
	routes.UserRoutes(app, *userService)

	// ============================================================ //
	// Start Server ...

	if config.Config().Environment == "develop" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
