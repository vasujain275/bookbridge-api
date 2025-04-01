package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vasujain275/bookbridge-api/docs"
	"github.com/vasujain275/bookbridge-api/internal/config"
	"github.com/vasujain275/bookbridge-api/internal/database"
	"github.com/vasujain275/bookbridge-api/internal/handler"
	"github.com/vasujain275/bookbridge-api/internal/middleware"
	"github.com/vasujain275/bookbridge-api/internal/repository"
	"github.com/vasujain275/bookbridge-api/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set Gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database connection
	db, err := database.New(cfg.PostgresConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	repo := repository.New(db.Pool)

	// Initialize services
	userService := service.NewUserService(repo)

	// Initialize router
	router := gin.Default()

	// Setup global middleware
	middleware.SetupGlobalMiddleware(router)

	// Register Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register user routes
	userHandler := handler.NewUserHandler(userService)
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/:id", userHandler.GetUser)       // GET /users/{id}
		userRoutes.GET("", userHandler.ListUsers)         // GET /users?limit=&offset=
		userRoutes.POST("", userHandler.CreateUser)       // POST /users
		userRoutes.PUT("/:id", userHandler.UpdateUser)    // PUT /users/{id}
		userRoutes.DELETE("/:id", userHandler.DeleteUser) // DELETE /users/{id}
	}

	// Create server
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
