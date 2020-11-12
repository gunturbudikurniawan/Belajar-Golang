package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// untuk ambil endpoint di controller
var route = controller.server()

// untuk ambil env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}
func run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values")
	}

	route.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// This is for testing, when done, do well to comment
	// seed.Load(server.DB)

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	fmt.Printf("Listening to port %s", apiPort)

	route.run(apiPort)
}
