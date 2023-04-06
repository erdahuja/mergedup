package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mergedup/business/data/dbschema"

	"mergedup/business/sys/database"
	"mergedup/foundation/config"

	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	cfg, err := config.LoadConfig("dev", ".")
	if err != nil {
		return err
	}

	dbConfig := cfg.GetDBConfig()

	if err := Migrate(dbConfig); err != nil {
		return err
	}

	if err := Seed(dbConfig); err != nil {
		return err
	}

	return nil
}

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")

// Migrate creates the schema in the database.
func Migrate(cfg config.DatabaseConfigurations) error {
	db, err := database.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	if err := dbschema.DropAll(context.TODO(), db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	if err := dbschema.Migrate(context.TODO(), db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")
	return nil
}

// Seed loads test data into the database.
func Seed(cfg config.DatabaseConfigurations) error {
	db, err := database.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := dbschema.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}
