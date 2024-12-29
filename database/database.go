package database

import (
	"embed"
	"errors"
	"fmt"
	"github.com/andReyM228/lib/log"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func InitDatabase(log log.Logger, config DBConfig, fs embed.FS) *gorm.DB {
	log.Debug("opening database connection")

	db, err := connect(config)
	if err != nil {
		log.Infof("main database connection failed: %s", err)
		return nil
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

func connect(config DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm.DB: %v", err)
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database failed: %v", err)
	}

	return db, nil
}

func dbAutoMigrate(db *gorm.DB, fs embed.FS, cfg DBConfig) (int, error) {
	migrate.SetTable("gorp_migrations")

	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: fs,
		Root:       cfg.ConfigDirectory,
	}

	dbSql, err := db.DB()
	if err != nil {
		return 0, errors.New(fmt.Sprintf("convert database: %v", err))
	}

	return migrate.Exec(dbSql, "postgres", migrations, migrate.Up)
}
