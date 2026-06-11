package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"youpiteron.dev/wildlands-backend/internal/application/repository"
	"youpiteron.dev/wildlands-backend/internal/application/service"
	"youpiteron.dev/wildlands-backend/internal/db"
	"youpiteron.dev/wildlands-backend/internal/env"
	"youpiteron.dev/wildlands-backend/internal/infrastructure"
	"youpiteron.dev/wildlands-backend/internal/infrastructure/postgres"
	"youpiteron.dev/wildlands-backend/internal/transport"
	"youpiteron.dev/wildlands-backend/internal/transport/controller"
)

const PORT = ":8080"

func main() {
	logger := infrastructure.NewConsoleLogger()
	err := env.LoadEnv()
	if err != nil {
		logger.Error("Failed to load env", slog.String("error", err.Error()))
		return
	}

	pool, err := db.NewPostgresPool(env.PostgresURL.GetValue())
	if err != nil {
		logger.Error("Failed to create postgres pool", slog.String("error", err.Error()))
		return
	}

	eventStore := postgres.NewEventStore(pool)
	snapshotStore := postgres.NewSnapshotStore(pool)
	matchMetadataStore := postgres.NewMatchMetadataStore(pool)
	playerStore := postgres.NewPlayerStore(pool)

	matchRepository := repository.NewMatchAggregate(snapshotStore, eventStore, matchMetadataStore, logger)
	matchService := service.NewMatch(
		eventStore,
		matchMetadataStore,
		playerStore,
		matchRepository,
		logger,
	)
	playerService := service.NewPlayer(playerStore, logger)

	matchController := controller.NewMatch(matchService, logger)
	playerController := controller.NewPlayer(playerService, logger)
	wsController := controller.NewWS(matchService, logger)

	router := transport.NewRouter(matchController, playerController, wsController)

	logger.Info(fmt.Sprintf("Starting server on port http://localhost%s", PORT))
	err = http.ListenAndServe(PORT, router)
	if err != nil {
		logger.Error("Failed to start server", slog.String("error", err.Error()))
		return
	}
}
