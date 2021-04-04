package api

import (
	"github.com/abhay676/auth/api/controllers"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}
}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Print("Not able to get ENV variables")
	}
	server.Initialize(os.Getenv("URI"), os.Getenv("DATABASE"))
	server.Run(":8000")
}