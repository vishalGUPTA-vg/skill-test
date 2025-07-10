package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	handlerPkg "go-service/handler"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file ")
	}

	router := handlerPkg.GetRouter()
	port := os.Getenv("PORT")
	log.Printf("Go pfd service running on port %s\n", port)
	router.Run(":" + port)
}
