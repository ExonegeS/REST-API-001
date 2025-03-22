package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	db "github.com/ExonegeS/REST-API-001/internal/adapter/postgres"
	"github.com/ExonegeS/REST-API-001/internal/api/http/handler"
	"github.com/ExonegeS/REST-API-001/internal/config"
	"github.com/ExonegeS/REST-API-001/internal/repository"
	"github.com/ExonegeS/REST-API-001/internal/service"
	"github.com/ExonegeS/REST-API-001/internal/usecase"
)

func main() {
	slog.Info(fmt.Sprintf("STARTING USERS SERVICE (%v environment) ...", os.Getenv("SERVICE_ENV")))

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("Error occured while loading config: %s", err))
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// TODO:
	// init DB connection
	dbConn, err := db.ConnectToPostgresDB(cfg.Database.HOST, cfg.Database.PORT, cfg.Database.USER, cfg.Database.PASS, cfg.Database.NAME)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer db.DisconnectFromPostgresDB(dbConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = ctx
	defer cancel()

	usersRepository := repository.NewUsersRepository(dbConn)
	usersUseCase := usecase.NewUsersUseCase(usersRepository)

	svc := service.NewUsersService(usersUseCase)
	svc = service.NewLoggingService(logger, svc)

	srv := handler.NewApiServer(svc)
	log.Fatal(srv.Start(cfg.Server.Port))
	// init server
	// Start server listening
	// Create gracefull shutdown
}
