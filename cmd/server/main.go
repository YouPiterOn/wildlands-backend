package main

import (
	"log/slog"
	"net/http"

	"youpiteron.dev/wildlands-backend/internal/application/repository"
	"youpiteron.dev/wildlands-backend/internal/application/service"
	"youpiteron.dev/wildlands-backend/internal/db"
	"youpiteron.dev/wildlands-backend/internal/env"
	"youpiteron.dev/wildlands-backend/internal/infrastructure"
	"youpiteron.dev/wildlands-backend/internal/infrastructure/postgres"
	"youpiteron.dev/wildlands-backend/internal/transport"
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

	matchRepository := repository.NewMatchAggregate(snapshotStore, eventStore, matchMetadataStore, logger)
	matchService := service.NewMatch(eventStore, matchMetadataStore, matchRepository, logger)

	wsHandler := transport.NewWSHandler(matchService)
	router := transport.NewRouter(wsHandler)

	logger.Info("Starting server on port http://localhost%s", PORT)
	err = http.ListenAndServe(PORT, router)
	if err != nil {
		logger.Error("Failed to start server", slog.String("error", err.Error()))
		return
	}
}
