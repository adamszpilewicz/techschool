package api

import (
	"github.com/gin-gonic/gin"
	db "techschool/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
