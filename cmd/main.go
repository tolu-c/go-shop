package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := api.run(api.mount())

	if err != nil {
		log.Printf("server has failed to start with error: %v", err)

		os.Exit(1)
	}

}
