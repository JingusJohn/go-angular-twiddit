package main

import (
	"os"

	"github.com/JingusJohn/go-angular-twiddit/backend/api"
	"github.com/JingusJohn/go-angular-twiddit/backend/config"
	"github.com/JingusJohn/go-angular-twiddit/backend/handlers"
	"github.com/JingusJohn/go-angular-twiddit/backend/storage"
)

func main() {
	config.LoadEnvironment()

	storage.ConnectToDB()
	storage.ConnectToRedis()

	app := api.SetupRouter()

	// add routes
	handlers.SetupHandlers(app)

	app.Listen(":" + os.Getenv("PORT"))
}
