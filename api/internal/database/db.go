package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*pgxpool.Pool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	dbURI := os.Getenv("DATABASE_URL")
	ctx := context.Background()

	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Minute * 20
	const defaultMaxConnIdleTime = time.Minute * 10
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	dbConfig, err := pgxpool.ParseConfig(dbURI + "?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to create a config to pool: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		log.Println("Erro ao criar o pool:", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		log.Printf("failed to ping database: %v", err)
		return nil, err
	}

	log.Println("Connected to database")
	return pool, nil
}
