package main

import (
	"log"
	"os"
	"time"

	"github.com/nguyenvantuan2391996/patient-order-number/handler"
	"github.com/nguyenvantuan2391996/patient-order-number/handler/middlewares"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/patient"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/infrastructure/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func initDatabase() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	return gorm.Open(mysql.Open(viper.GetString("DB_SOURCE")), &gorm.Config{
		Logger: newLogger,
	})
}

func main() {
	viper.AddConfigPath("build")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	db, err := initDatabase()
	if err != nil {
		logrus.Fatal("failed to open database:", err)
		return
	}

	// repository
	accountRepo := repository.NewAccountRepository(db)

	// service
	patientService := patient.NewPatientService(accountRepo)

	h := handler.NewHandler(patientService)

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middlewares.Recover())

	// public
	v1API := r.Group("v1")
	{
		v1API.Use(middlewares.APIKeyAuthentication())

		v1API.POST("/test", h.Test)
	}

	err = r.Run(":" + viper.GetString("PORT"))
	if err != nil {
		return
	}
}
