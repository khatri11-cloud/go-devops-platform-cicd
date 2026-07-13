package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"devops-platform/internal/config"
	"devops-platform/internal/handlers"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to DevOps Platform")
}

func main() {
	cfg := config.Load()
	http.HandleFunc("/", home)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/version", handlers.VersionHandler(cfg.AppVersion))

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: nil,
	}

	go func() {
		log.Printf("Starting server on :%s", cfg.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped gracefully")
}
