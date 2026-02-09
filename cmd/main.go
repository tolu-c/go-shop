package main

import (
	"log"
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

	err := api.run(api.mount())

	if err != nil {
		log.Printf("server has failed to start with error: %v", err)

		os.Exit(1)
	}

}
