package handlers

import (
	"github.com/JingusJohn/go-angular-twiddit/backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupHandlers(app *fiber.App) {
	protected := app.Group("/api/protected", middleware.NewMiddleware(middleware.AuthRequired))
	// auth routes
	app.Post("/api/users", CreateUser)
	app.Post("/api/login", Login)
	protected.Post("/logout", Logout)
	protected.Get("/auth/health", HealthCheck)

	// post routes
	protected.Post("/posts", CreatePost)
	app.Get("/api/posts", GetPosts)
	protected.Post("/posts/:id", DeletePost)
	protected.Post("/posts/:id/rate", RatePost)
	app.Get("/api/posts/:id", GetPost)
	protected.Post("/posts/:id/comments", CommentOnPost)
	protected.Delete("/posts/:id/comments/:comment_id", DeleteComment)
}
