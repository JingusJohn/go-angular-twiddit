package api

import (
	"os"
	"time"

	"github.com/JingusJohn/go-angular-twiddit/backend/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	// SessionStore is the global session store
	SessionStore *session.Store
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	config := session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 5,
		CookieSecure:   true,
		Storage:        storage.RedisStore,
	}

	// if we're in DEV, turn https off
	if os.Getenv("ENV") == "DEV" {
		config.CookieSecure = false
	}

	sessionStore := session.New(config)
	SessionStore = sessionStore

	//app.Use(sessionStore)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
	}))

	return app
}
