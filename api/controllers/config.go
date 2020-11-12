package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gunturbudikurniawan/Belajar-Golang/api/middlewares"
	"github.com/gunturbudikurniawan/Belajar-Golang/api/models"

	// for postgre
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Server for blala
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var errList = make(map[string]string)

//Initialize display the claims licely in the terminal
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error connecting to postgres:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	} else {
		fmt.Println("Unknown Driver")
	}
	server.DB.Debug().AutoMigrate(
		&models.Admin{},
		&models.Saved_orders{},
		&models.Outlets{},
		&models.Sales{},
		&models.Onlinesales{},
	)
	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())
	server.initialRoutes()
}

// Run for read server
func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
