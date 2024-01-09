package main

import (
	"hello/pkg/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users/:userId", server.GetUserByID)
	router.Run(":9001")

	// user, err := client.GetUserByID("localhost:8083", "1")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)

}
