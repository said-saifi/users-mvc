package app

import (
	"github.com/said-saifi/users-mvc/controllers"
)

func mapUrls(){
	router.GET("/heartbeat", controllers.Heartbeat)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/users/search", controllers.SearchUser)
}