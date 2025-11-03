package main

import (
	"log"

	"github.com/aok207/go-social/internal/db"
	"github.com/aok207/go-social/internal/env"
	"github.com/aok207/go-social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("PORT", ":6969"),
		db: dbConfig{
			addr:         env.GetString("DB_CONN_STRING", "postgres://user:********@localhost:5432/go_social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS",
				30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.NewDB(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	log.Println("DB connection established")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Fatal(app.serve(app.registerRoutes()))
}
