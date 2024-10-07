package routes

import (
	"gobackend/controllers"
	"gobackend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/admin/register", controllers.Register)
	app.Post("/admin/login", controllers.Login)
	app.Static("/api/uploads", "./uploads")
	//mobile
	app.Use("api", middlewares.IsMobileJWT)
	app.Post("/api/words/:id", controllers.WordsByTopicId)
	app.Post("/api/rand-words/:id", controllers.RandomWordsByTopicId)
	app.Post("/api/topics", controllers.Topics)
	//admin
	app.Use("admin", middlewares.IsAuthenticated)
	app.Get("/admin/user", controllers.User)
	app.Post("/admin/logout", controllers.Logout)

	app.Get("/admin/topics", controllers.AllTopics)
	app.Post("/admin/topics", controllers.CreateTopic)
	app.Get("/admin/topics/:id", controllers.GetTopic)
	app.Put("/admin/topics/:id", controllers.UpdateTopic)
	app.Delete("/admin/topics/:id", controllers.DeleteTopic)

	app.Get("/admin/words", controllers.AllWords)
	app.Post("/admin/words", controllers.CreateWord)
	app.Get("/admin/words/:id", controllers.GetWord)
	app.Put("/admin/words/:id", controllers.UpdateWord)

	app.Post("/admin/image/upload", controllers.UploadImage)
	app.Post("/admin/audio/upload", controllers.UploadAudio)

	// app.Put("/api/users/info", controllers.UpdateInfo)
	// app.Put("/api/users/password", controllers.UpdatePassword)

	// app.Get("/api/users", controllers.AllUsers)
	// app.Post("/api/users", controllers.CreateUser)
	// app.Get("/api/users/:id", controllers.GetUser)
	// app.Put("/api/users/:id", controllers.UpdateUser)
	// app.Delete("/api/users/:id", controllers.DeleteUser)

	// app.Get("/api/roles", controllers.AllRoles)
	// app.Post("/api/roles", controllers.CreateRole)
	// app.Get("/api/roles/:id", controllers.GetRole)
	// app.Put("/api/roles/:id", controllers.UpdateRole)
	// app.Delete("/api/roles/:id", controllers.DeleteRole)

	// app.Get("/api/permissions", controllers.AllPermissions)

	// app.Get("/api/words/:id", controllers.GetWord)
	// app.Put("/api/words/:id", controllers.UpdateWord)
	// app.Delete("/api/words/:id", controllers.DeleteWord)

	// app.Get("/api/orders", controllers.AllOrders)
	// app.Post("/api/export", controllers.Export)
	// app.Get("/api/chart", controllers.Chart)
}
