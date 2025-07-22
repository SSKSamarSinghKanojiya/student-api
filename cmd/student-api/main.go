package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SSKSamarSinghKanojiya/student-api/internal/config"
	"github.com/SSKSamarSinghKanojiya/student-api/internal/http/student"
	"github.com/SSKSamarSinghKanojiya/student-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/student", student.New(storage))
	router.HandleFunc("GET /api/student/{id}", student.GetById(storage))

	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started ", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start error")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server ", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfulyl")
}

// go run cmd/student-api/main.go -config config/local.yaml
