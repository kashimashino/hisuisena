package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kashimashino/hackweek/api/v1"
	"github.com/kashimashino/hackweek/middleware"
	"github.com/kashimashino/hackweek/model"
	"github.com/kashimashino/hackweek/server"
)



func InitRouter()  {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(200, "index.html", nil)
	//})
	r.Use(middleware.CORS())
	r.POST("login", server.Login)
	router := r.Group("user")
	figure := r.Group("figure")
	//用户模块

	router.Use(middleware.JwtToken())
	{

		router.GET("/",v1.GetUser)
		router.PUT("/editInfo/:id",v1.EditUser)
		router.PUT("/editPwd/:id",v1.EditPassword)
		router.DELETE("/:id",v1.DeleteUser)
		router.POST("/makeActivity",v1.MakeActivity)
		router.GET("checkActivities",v1.CheckActivities)

	}
	r.POST("user/add",v1.AddUser)
	//r.POST("/upload",v1.Upload)
	r.POST("/search",v1.Search)

	//功能模块

	figure.Use(middleware.JwtToken())
	figure.GET("/:id", model.CheckClothes)

  	_ = r.Run(":8080")
}
