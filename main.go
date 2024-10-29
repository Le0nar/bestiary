package main

import (
	"fmt"

	"github.com/Le0nar/bestiary/internal/repository"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := repository.NewPostgresDB(repository.DatabaseConfig{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	fmt.Printf("db: %v\n", db)
}