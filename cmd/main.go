package main

import (
	"context"
	"log"
	"os"

	"github.com/Pyth0nHater/link-shorter/internal/database"
	"github.com/Pyth0nHater/link-shorter/internal/handler"
	"github.com/Pyth0nHater/link-shorter/internal/repository"
	"github.com/Pyth0nHater/link-shorter/internal/service"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbType := os.Getenv("DB_TYPE")
	dbUrl := os.Getenv("DB_URL")

	ctx := context.Background()	

	var linkRepo repository.LinkRepositoryInterface

	switch (dbType){
		case "postgres":
			db := database.InitDB(ctx, dbUrl)
			linkRepo = repository.NewPostgresLinkRepository(db)
		case "memory":
			mem:= database.InitMemory()
			linkRepo = repository.NewMemeoryLinkRepository(mem)
		default:
			mem:= database.InitMemory()
			linkRepo = repository.NewMemeoryLinkRepository(mem)
	}

	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	router := gin.Default()

	api := router.Group("/api/v1")
	{

		links := api.Group("/link")
		{
			links.POST("/create", linkHandler.CreateLink)
			links.GET("/get", linkHandler.GetLink)
		}
	}
	router.Run()
}
