package main

import (
	"context"
	"log"
	"os"
	"url-shorting/cmd/api/handler"
	"url-shorting/cmd/internal"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up Gin router and dependencies

	ctx := context.Background()

	router := gin.Default()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := os.Getenv("MONGO_URI")
	mongoCLinet, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
		panic(err)
	}

	repo := internal.NewRepository(mongoCLinet)
	svc := internal.NewService(repo)
	h := handler.NewHandler(svc)

	router.POST("/URLs", h.GenerateURL)

	router.Run("localhost:8080")
}
