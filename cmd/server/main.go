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

	http.ListenAndServe(":8080", router)
}
