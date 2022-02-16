package main

import (
	"fmt"
	"github.com/uptrace/bunrouter"
	"net/http"
	"src/pkg/handle"
)

func main() {

	router := bunrouter.New()
	router.GET("/", handle.GetTwitters)
	router.POST("/tweet", handle.CreateTweet)

	fmt.Println("Server running...")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}

}
