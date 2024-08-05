package main

import (
	"fmt"
	"os"
	handlers "profileService/internal/handlers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {

	fmt.Println(os.Getenv("LOL"))
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		create := v1.Group("/create_profile")
		{
			create.POST("", handlers.CreateProfile)
		}
		read := v1.Group("/get_profile")
		{
			read.POST("", handlers.GetProfile)
		}
		update := v1.Group("/update_profile")
		{
			update.GET("", handlers.UpdateProfile)
		}
		delete := v1.Group("/delete_profile")
		{
			delete.GET("", handlers.DeleteProfile)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
