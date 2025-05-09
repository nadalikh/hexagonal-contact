package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"graph/domain"
	"graph/logger"
	"graph/service"
)

func Start() {
	systemConfig := NewSystemConfig()
	if systemConfig.Gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	db := getDbClient(
		systemConfig.PostgresConfig.Username,
		systemConfig.PostgresConfig.Password,
		systemConfig.PostgresConfig.Server,
		systemConfig.PostgresConfig.Port,
		systemConfig.PostgresConfig.DBName)

	contactHandler := NewContactRestHandler(service.NewContactService(domain.NewContactPostgresRepo(db)))
	phoneNumberHandler := NewPhoneNumberRestHandler(service.NewPhoneNumberService(domain.NewPhoneNumberPostgresRepo(db)))

	router.POST("/contact", contactHandler.CreateOne)
	router.GET("/contact/search", contactHandler.Search)
	router.PUT("/contact/update", contactHandler.Update)
	router.POST("/phone-number/add", phoneNumberHandler.AddToContact)

	err := router.Run(fmt.Sprintf("%s:%s", systemConfig.ServerConfig.Server, systemConfig.ServerConfig.Port))
	if err != nil {
		panic(err)
	}
}

func getDbClient(username, passwd, host, port, dbName string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", host, username, passwd, dbName, port)), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}
	return db
}
