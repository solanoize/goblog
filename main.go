package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/solanoize/goblog/internal/controllers"
	"github.com/solanoize/goblog/internal/middlewares"
	"github.com/solanoize/goblog/internal/models"
	"github.com/solanoize/goblog/internal/repositories"
	"github.com/solanoize/goblog/internal/routers"
	"github.com/solanoize/goblog/internal/usecases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetLogger() *log.Logger {
	// Ganti nama file biar lebih umum
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 1. Prefix dikosongin aja ""
	// 2. Tambahin io.MultiWriter biar log muncul di terminal DAN file sekaligus
	multiWriter := io.MultiWriter(os.Stdout, file)

	return log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
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
	var port string = os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5000" // Default port kalau env kosong
	}
	var db *gorm.DB = GetDB()
	var logger *log.Logger = GetLogger()
	var mainRouter *chi.Mux = chi.NewRouter()

	mainRouter.Use(middleware.CleanPath)    // Otomatis ngerapiin // jadi /
	mainRouter.Use(middleware.StripSlashes) // /posts/ jadi /posts otomatis

	if err != nil {
		logger.Fatal(".env file not found, using system env")
	}

	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		logger.Fatal("Failed to migrate database: ", err)
	}

	// var userRouter routers.UserRouter = routers.NewUserRouter(db, logger)
	// router.Mount("/", userRouter.Register())

	// USER DEPENDENCY INJECTION
	var userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	var userUseCase usecases.UserUseCase = usecases.NewUserUseCase(logger, userRepository)
	var userController controllers.UserController = controllers.NewUserController(userUseCase)
	var authMiddleware middlewares.AuthMiddleware = middlewares.NewAuthMiddleware(logger, userUseCase)

	// POST DEPENDENCY INJECTION
	var postRepository repositories.PostRepository = repositories.NewPostRespository(db)
	var postUseCase usecases.PostUseCase = usecases.NewPostUseCase(logger, postRepository)
	var postController controllers.PostController = controllers.NewPostController(postUseCase)

	routers.NewUserRouter(mainRouter, authMiddleware, userController).Register()
	routers.NewPostRouter(mainRouter, authMiddleware, postController).Register()

	logger.Printf("Server running on port %s\n", port)
	if err = http.ListenAndServe(":"+port, mainRouter); err != nil {
		logger.Fatal("Server failed to start: ", err)
	}
}
