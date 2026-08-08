package main

func main() {
	cfg := config{
		port: ":8080",
		db: dbConfig{},
	}

	api := app{
		config: cfg,
	}
}
