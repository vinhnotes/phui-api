package main

import (
	"bongdaphui/api"
	"bongdaphui/database"
	"bongdaphui/lib/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "bongdaphui/docs"
)

// @title Phủi API
// @version 1.0
// @description Bóng đá phủi social network.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /v1
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default()                                            // create gin app
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port
}
