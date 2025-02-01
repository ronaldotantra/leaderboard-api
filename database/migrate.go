package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/mattes/migrate/source/file"
	"github.com/ronaldotantra/leaderboard-api/config"
	"github.com/ronaldotantra/leaderboard-api/internal/logger"
	"github.com/ronaldotantra/leaderboard-api/internal/logger/logrus"
)

func main() {
	downFlag := flag.Bool("down", false, "database migration down")
	envFlag := flag.String("envFile", "", "target env file")
	flag.Parse()

	if envFlag == nil || *envFlag == "" {
		log.Fatal("envFile flag not provided")
		return
	}

	err := godotenv.Load(*envFlag)
	if err != nil {
		log.Fatal("error: failed to load the env file")
		return
	}

	config.Init()
	level := logger.Debug
	if config.IsProductionEnvironment() {
		level = logger.Info
	}

	// Initialize of Logger
	logConfig := &logger.Configuration{
		ConsoleJSONFormat: true,
		ConsoleLevel:      level,
	}

	logger.SetRepository(logrus.NewLogrusLogger(logConfig))

	logger.Infof("Opening database connection")

	db, err := sql.Open("postgres", config.DatabaseConnectionString)
	logger.Infof("Database connected.")
	if err != nil {
		logger.Fatalf("error opening migration database - %v\n", err)
		return
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatalf("error generating postgres instance - %v\n", err)
		return
	}
	logger.Infof("Postgres instance generated.")

	logger.Infof("Opening migration files...")
	fsrc, err := (&file.File{}).Open("file://database/migrations")
	if err != nil {
		logger.Fatalf("error opening migration files - %v", err)
		return
	}
	logger.Infof("Migration files opened.")

	logger.Infof("Creating migration instance...")
	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		logger.Fatalf("error generating migrate instance - %v", err)
		return
	}
	logger.Infof("Migration instance created.")

	if *downFlag {
		logger.Infof("Rollback migration..")
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			logger.Fatalf("error rollback migrations - %v", err)
			return
		}
		version, _, _ := m.Version()
		logger.Infof("Rollback complete to version %d.\n", version)
	} else {
		logger.Infof("Migrating migration..")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			logger.Fatalf("error migrating migrations - %v", err)
			return
		}
		logger.Infof("Migrate complete.")
	}
}
