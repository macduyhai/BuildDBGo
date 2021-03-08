package main

import (
	"BuildDBGo/controllers"
	"BuildDBGo/daos"
	"BuildDBGo/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CORSMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitGin(db *gorm.DB) *gin.Engine {
	userDao := daos.NewUserDao(db)
	userService := services.NewUserService(userDao)

	authenDao := daos.NewAuthenDao(db)
	authenService := services.NewAuthenService(authenDao)

	ctl := controllers.Controller{
		AdScreenService: userService,
		AuthenService:   authenService,
	}

	engine := gin.Default()
	engine.Use(CORSMiddleWare())

	engine.GET("/health", ctl.HealthCheck)
	apiGroup := engine.Group("/api/v1")
	{
		authenGroup := apiGroup.Group("/authenkey")
		{
			authenGroup.POST("", ctl.AddAuthen)
			authenGroup.GET("", ctl.SearchAuthen)
		}
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("", ctl.AddUser)
			// orderGroup.DELETE("", ctl.Delete)
			// orderGroup.PUT("", ctl.UpdateOrders)
			// orderGroup.GET("/search", ctl.Search)
		}

	}
	return engine
}
