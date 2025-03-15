package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"productManagement/internal/controller"
	"productManagement/internal/repository"
	"productManagement/internal/usecase"
	"productManagement/pkg/dbsvc"
	"productManagement/pkg/logger"
	"productManagement/route"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := dbsvc.StartDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	ctx := context.Background()
	// Initialize logger with context (this now includes trace/span IDs)
	logger.Init(ctx)
	e := gin.Default()

	fmt.Println("Successfully connected to the database")
	repo := repository.NewRepository(dbsvc.GetDBConn())
	usecase := usecase.NewUseCase(repo)
	controller := controller.NewController(usecase)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true

	e.Use(cors.New(corsConfig))
	route.InitRoutes(e, controller)

	port := "8080"
	srv := &http.Server{
		Addr:    "localhost:" + port,
		Handler: e,
	}

	// Run the server in a separate goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	logger.Info(fmt.Sprintf("SERVER LISTENING ON PORT %s", port))
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	logger.Info("Server exiting")
}
