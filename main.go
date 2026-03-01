package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/solanoize/goblog/internal/models"
	"github.com/solanoize/goblog/internal/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetLogger() *log.Logger {
	var file *os.File
	var err error

	file, err = os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	return log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetDB() *gorm.DB {
	var dsn string = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	var err error
	var db *gorm.DB

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB: ", err)
	}

	// SET CONNECTION POOL DI SINI
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Database ready")

	return db
}

func main() {
	var err error
	err = godotenv.Load()
	var db *gorm.DB = GetDB()
	var logger *log.Logger = GetLogger()
	var router *chi.Mux = chi.NewRouter()

	if err != nil {
		logger.Fatal(".env file not found, using system env")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Fatal("Failed to migrate database: ", err)
	}

	var userRouter routers.UserRouter = routers.NewUserRouter(db, logger)
	router.Mount("/", userRouter.Register())

	fmt.Printf("Server running on port %s\n", os.Getenv("SERVER_PORT"))
	http.ListenAndServe(os.Getenv("SERVER_PORT"), router)
}
