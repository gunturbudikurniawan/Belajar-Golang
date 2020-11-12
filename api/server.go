package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gunturbudikurniawan/Belajar-Golang/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

// Run for babab
func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	apiPort := os.Getenv("PORT")
	fmt.Printf("Listening to port %s", "4001")
	if apiPort == "" {
		apiPort = "4001"
	}
	server.Run(":" + apiPort)
}
