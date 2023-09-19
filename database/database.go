package database

import (
	"embed"
	"errors"
	"fmt"
	"github.com/andReyM228/lib/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

type DBConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	DBName          string `yaml:"db-name"`
	AutoMigrate     bool   `yaml:"auto-migrate"`
	ConfigDirectory string `yaml:"config-directory"`
}

func InitDatabase(log log.Logger, config DBConfig, fs embed.FS) *sqlx.DB {
	log.Debug("opening database connection")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}

	if err := db.Ping(); err != nil {
		db, err = createDataBase(log, config)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if config.AutoMigrate {
		countMigration, err := dbAutoMigrate(db, fs, config)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Infof("migration applied: %d", countMigration)
	}

	return db
}

func createDataBase(log log.Logger, config DBConfig) (*sqlx.DB, error) {
	log.Debug("opening database connection")

	newDBName := config.DBName

	config.DBName = "postgres"

	db, err := connect(config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("default db: %v", err))
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", newDBName))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("create new db: %v", err))
	}

	config.DBName = newDBName

	db, err = connect(config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("new db: %v", err))
	}

	return db, nil
}

func connect(config DBConfig) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open database: %v", err))
	}

	if err := db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("ping database: %v", err))
	}

	return db, nil
}

func dbAutoMigrate(db *sqlx.DB, fs embed.FS, cfg DBConfig) (int, error) {
	migrate.SetTable("gorp_migrations")

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: fs,
		Root:       cfg.ConfigDirectory,
	}

	return migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
}
