package main

import (
	"log"
	"os"
	"time"

	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	"github.com/nguyenvantuan2391996/patient-order-number/handler"
	"github.com/nguyenvantuan2391996/patient-order-number/handler/middlewares"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/admin"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/auth"
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
	patientRepo := repository.NewPatientRepository(db)

	// service
	patientService := patient.NewPatientService(accountRepo, patientRepo)
	adminService := admin.NewAdminService(accountRepo)
	authService := auth.NewAuthService(accountRepo)

	h := handler.NewHandler(patientService, adminService, authService)

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

	// Load HTML templates from the "templates" folder
	r.LoadHTMLGlob("templates/*")

	// Admis APIs
	adminAPI := r.Group("v1/api")
	{
		// auth
		adminAPI.Use(middlewares.JWTValidationMW(constants.RoleAdmin))

		adminAPI.POST("/accounts", h.CreateAccount)
		adminAPI.PUT("/accounts/:user_id", h.UpdateAccount)
		adminAPI.DELETE("/accounts/:user_id", h.DeleteAccount)
	}

	// Public APIs
	v1Public := r.Group("v1")
	{
		v1Public.GET("/patient-page", h.GetPatientPage)
		v1Public.GET("/patient/login", h.LoginPatientPage)
		v1Public.GET("/patient/admin", h.GetAdminPatientPage)
		v1Public.GET("/patient/:channel", h.InitWSPatient)
	}

	// Patient APIs
	v1Patient := r.Group("v1/api")
	{
		v1Patient.Use(middlewares.JWTValidationMW(constants.RoleAdmin))

		v1Patient.POST("/patient", h.CreatePatient)
		v1Patient.GET("/patient/list", h.GetListPatient)
		v1Patient.PUT("/patient/:id", h.UpdatePatient)
		v1Patient.DELETE("/patient/:id", h.DeletePatient)
	}

	err = r.Run(":" + viper.GetString("PORT"))
	if err != nil {
		return
	}
}
