package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	healthhandler "crmata-go/internal/adapters/http/handler/health"
	usershandler "crmata-go/internal/adapters/http/handler/users"
	httproutes "crmata-go/internal/adapters/http/routes"
	postgresadapter "crmata-go/internal/adapters/postgres"
	usersservice "crmata-go/internal/application/service/users"
	"crmata-go/internal/infrastructure/config"
	"crmata-go/internal/infrastructure/database"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pool, err := database.Connect(ctx, cfg.PostgresDSN, 15, 2*time.Second)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer pool.Close()

	if err := database.RunMigrations(ctx, pool); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	repo := postgresadapter.NewUserRepository(pool)
	service := usersservice.NewService(repo)
	usersHandler := usershandler.New(service)
	healthHandler := healthhandler.New()
	router := httproutes.NewRouter(healthHandler, usersHandler)

	server := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("API listening on :%s", cfg.AppPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}
