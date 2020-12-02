package app

import (
	"github.com/mohsanabbas/ticketing_users_api/controllers/ping"
	"github.com/mohsanabbas/ticketing_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)

}
