package main

import (
	"log"
	"simpleApic/config"
	personRepo "simpleApic/person/repo"
	personUsecase "simpleApic/person/usecase"
	personHandler "simpleApic/person/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := "7861"

	db := config.DbConnect()
	defer db.Close()

	router := gin.Default()

	personRepo := personRepo.CreatePersonRepo(db)
	personUsecase := personUsecase.CreatePersonUsecase(personRepo)
	personHandler.CreatePersonHandler(router, personUsecase)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
