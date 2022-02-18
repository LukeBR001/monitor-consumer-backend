package main

import (
	"fmt"
	"github.com/uptrace/bunrouter"
	"net/http"
	"src/pkg/database"
	"src/pkg/handle"
	"src/pkg/repository"
	"src/pkg/service"
)

func main() {

	dbConecct := database.CreateDbConnection()

	repo := repository.NewRepository(dbConecct)
	serv := service.NewService(repo)
	hand := handle.NewHandle(serv)

	router := bunrouter.New()
	router.GET("/", hand.GetTwitters)
	router.POST("/tweet", hand.CreateTweet)

	fmt.Println("Server running...")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}

}
