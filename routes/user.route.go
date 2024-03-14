package routes

import (
	"github.com/mesxx/Fiber_User_Management_API/handlers"
	"github.com/mesxx/Fiber_User_Management_API/middlewares"
	"github.com/mesxx/Fiber_User_Management_API/repositories"
	"github.com/mesxx/Fiber_User_Management_API/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewUserRepositoy(db)
	usecase := usecases.NewUserUsecase(repository)
	handler := handlers.NewUserHandler(usecase)

	router.Post("/", handler.Create)
	router.Post("/login", handler.Login)
	router.Get("/", handler.GetAll)
	router.Post("/password/forgot", handler.ForgotPassword)
	router.Patch("/password/reset", handler.ResetPassword)
	router.Delete("/delete/all", handler.DeleteAll)

	// Authorization
	router.Use(middlewares.RestrictedUser)

	router.Get("/account", handler.GetByID)
	router.Patch("/account", handler.UpdateByID)
	router.Delete("/account", handler.DeleteByID)
	//
}
