package app

import (
	"github.com/eremitic/bookstore_users-api/controllers/ping"
	"github.com/eremitic/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:id", users.GetUser)
	router.POST("/users", users.CreateUser)
	//router.GET("/users/search",controllers.SearchUser)
	//123
}
