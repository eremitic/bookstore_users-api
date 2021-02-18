package app

import (
	"github.com/eremitic/bookstore_users-api/controllers/ping"
	"github.com/eremitic/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:id", users.Get)
	router.PUT("/users/:id", users.Update)
	router.POST("/users", users.Create)
	router.DELETE("/users/:id", users.Delete)
	router.GET("/internal/users/search", users.Search)

}
