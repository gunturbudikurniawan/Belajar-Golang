package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gunturbudikurniawan/Belajar-Golang/api/middlewares"
)

func (s *Server) initialRoutes() {
	v1 := s.Router.Group("/api")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	v2 := s.Router.Group("/api/admin")
	{
		v2.POST("/login", s.Login)
		v2.POST("/register", s.CreateAdmin)
		v2.GET("/getAdmin", s.GetAdmins)
		v2.GET("/getuser/:id", s.GetUser)
		v2.PUT("/updateuser/:id", middlewares.TokenAuthMiddleware(), s.UpdateUser)
		v1.DELETE("/deleteuser/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)
		v1.GET("/ShowSleep", middlewares.TokenAuthMiddleware(), s.GetShow)
	}
}
