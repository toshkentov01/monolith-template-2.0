package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/toshkentov01/template/api/controllers"
	service "github.com/toshkentov01/template/service/user"
)

// UserRoutes func for describe group of public routes.
func UserRoutes(app *fiber.App, userService service.UserService) {
	route := app.Group("/api/v1")
	api := controllers.NewAPI(userService)

	// Routes For Post Method:
	route.Post("/register/signup/", api.SignUp)

	// Routes For Get Method:
	route.Get("/user/profile/:user_id/", api.GetProfile)
	route.Get("/user/get-my-profile/", api.GetMyProfile)

}
