package configs

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Env    string `default:"dev" envconfig:"ENV"`
	Host   string `default:"0.0.0.0:3000" envconfig:"HOST"`
	DB     DbConfig
	Engine *gin.Engine
	Routes *gin.RouterGroup
}

func NewServer() (*Server, error) {

	dbConfig, err := LoadDbConfig()
	if err != nil {
		return nil, err
	}
	if dbConfig == nil {
		return nil, errors.New("cannot load database config")
	}
	server := Server{
		Env:    os.Getenv("ENV"),
		Host:   os.Getenv("HOST"),
		DB:     *dbConfig,
		Engine: gin.Default(),
		Routes: nil,
	}
	return &server, nil
}

func (s *Server) Listen() error {
	return s.Engine.Run()
}
