package handlers

import (
	"profileService/internal/db"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller struct {
	db *db.DataBase
}

func GetController() Controller {
	return Controller{db: db.InitDb()}
}

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
func (controller *Controller) Run() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		create := v1.Group("/create_profile")
		{
			create.POST("", controller.CreateProfile)
		}
		read := v1.Group("/get_profiles")
		{
			read.GET("", controller.GetProfiles)
		}
		update := v1.Group("/update_profile")
		{
			update.GET("", controller.UpdateProfile)
		}
		delete := v1.Group("/delete_profile")
		{
			delete.GET("", controller.DeleteProfile)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
