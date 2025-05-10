package app

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"graph/internal"
	"graph/logger"
	"os"
)

type SystemConfig struct {
	ServerConfig   ServerConfig
	PostgresConfig Postgress
	Gin            GinConfig
}
type GinConfig struct {
	ReleaseMode bool
}
type ServerConfig struct {
	Server string
	Port   string
}
type Postgress struct {
	Server   string
	Port     string
	DBName   string
	Username string
	Password string
}

func NewSystemConfig() *SystemConfig {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	systemConfig := new(SystemConfig)
	systemConfig.ServerConfig = ServerConfig{
		Server: os.Getenv("SERVER"),
		Port:   os.Getenv("PORT"),
	}
	systemConfig.PostgresConfig = Postgress{
		Server:   os.Getenv("POSTGRES_SERVER"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
	systemConfig.Gin = GinConfig{
		ReleaseMode: os.Getenv("GIN_MODE") == "release",
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("phone", internal.PhoneValidation); err != nil {
			logger.Error(err.Error())
		}
	}
	return systemConfig
}
