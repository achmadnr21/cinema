package main

import (
	"database/sql"
	"fmt"

	"github.com/achmadnr21/cinema/config"
	pgsql "github.com/achmadnr21/cinema/infrastructure/rdbms"
	"github.com/achmadnr21/cinema/internal/handler"
	"github.com/achmadnr21/cinema/internal/middleware"
	"github.com/achmadnr21/cinema/internal/repository"
	"github.com/achmadnr21/cinema/internal/usecase"
	"github.com/achmadnr21/cinema/internal/utils"
	gin_api "github.com/achmadnr21/cinema/service"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("=== Starting Application ===")
	envload, err := service_init()
	if err != nil {
		fmt.Printf("Error initializing service: %v\n", err)
		return
	}
	// load configuration
	var db *sql.DB = pgsql.GetPG()
	defer db.Close()

	var api gin_api.RESTapi
	var apiV *gin.RouterGroup = api.Init(gin.ReleaseMode)
	if apiV == nil {
		fmt.Println("[Error] API initialization failed")
		panic("API initialization failed")
	} else {
		fmt.Println("[Info] API initialized successfully")
	}

	// Initialize Middleware
	middleware.InitDbMiddleware(db)

	// Initialize REPO
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	// Initialize USECASE
	authUsecase := usecase.NewAuthUsecase(userRepo)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)

	// Initialize HANDLER
	handler.NewAuthHandler(apiV, authUsecase)
	handler.NewRoleHandler(apiV, roleUsecase)

	apiV.GET("/ping", HandlePing)

	fmt.Println("[Info] Starting server on port", envload.Port)
	api_service := fmt.Sprintf(":%d", envload.Port)
	api.Router.Run(api_service)
}

func service_init() (*config.Config, error) {
	// load configuration
	envload, _ := config.LoadConfig()
	if envload == nil {
		return &config.Config{}, fmt.Errorf("failed to load configuration")
	}

	err := pgsql.InitPG(envload.Database)
	if err != nil {
		return envload, fmt.Errorf("error initializing database: %w", err)
	}
	return envload, nil
}

func HandlePing(c *gin.Context) {
	c.JSON(200, utils.ResponseSuccess("Pong", &struct {
		Developer string `json:"developer"`
	}{
		Developer: "Achmad Nashruddin Riskynanda",
	}))
}
