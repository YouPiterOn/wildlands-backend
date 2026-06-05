package main

import (
	"log/slog"
	"net/http"

	"youpiteron.dev/wildlands-backend/internal/application"
	"youpiteron.dev/wildlands-backend/internal/db"
	"youpiteron.dev/wildlands-backend/internal/env"
	"youpiteron.dev/wildlands-backend/internal/infrastructure"
	"youpiteron.dev/wildlands-backend/internal/infrastructure/eventstore"
	"youpiteron.dev/wildlands-backend/internal/infrastructure/snapshotstore"
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

	eventStore := eventstore.NewPostgresEventStore(pool)
	snapshotStore := snapshotstore.NewPostgresSnapshotStore(pool)

	matchRepository := application.NewAggregateMatchRepository(snapshotStore, eventStore, logger)
	matchService := application.NewMatchService(eventStore, matchRepository, logger)

	wsHandler := transport.NewWSHandler(matchService)
	router := transport.NewRouter(wsHandler)

	logger.Info("Starting server on port http://localhost%s", PORT)
	err = http.ListenAndServe(PORT, router)
	if err != nil {
		logger.Error("Failed to start server", slog.String("error", err.Error()))
		return
	}
}
