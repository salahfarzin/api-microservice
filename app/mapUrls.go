package app

import "github.com/salahfarzin/api-microservice/handlers"

func mapUrls() {
	router.POST("/users", handlers.UsersHandler.Create)
	router.GET("/users/:id", handlers.UsersHandler.Get)
}
