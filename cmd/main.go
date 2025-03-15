package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		panic("could not load .env file")
	}
	port := os.Getenv("PORT")
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	server.Run(fmt.Sprintf(":%s", port))
}
