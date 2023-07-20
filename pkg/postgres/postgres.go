package postgres

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/slog"
	"time"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	// gomigrate migration resolver
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	// db driver
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type (
	Config struct {
		Postgres
	}

	Postgres struct {
		ConnString      string        `validate:"required"`
		MaxOpenConns    int           `validate:"required"`
		ConnMaxLifetime time.Duration `validate:"required"`
		MaxIdleConns    int           `validate:"required"`
		ConnMaxIdleTime time.Duration `validate:"required"`
		MigrationsPath  string        `validate:"required"`
		DBName          string        `validate:"required"`
		AutoMigrate     bool
	}
)

func InitPsqlDB(cfg Config, l *slog.SugaredLogger) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.Postgres.ConnString)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime * time.Second)

	err = db.Ping()
	if err != nil {
		l.Error(err)
		return nil, err
	}

	if cfg.Postgres.AutoMigrate {
		migrationDriver, err := postgres.WithInstance(db.DB, &postgres.Config{})
		if err != nil {
			l.Error(err)
			return nil, err
		}

		m, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", cfg.Postgres.MigrationsPath),
			cfg.Postgres.DBName,
			migrationDriver,
		)
		if err != nil {
			l.Error(err)
			return nil, err
		}

		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			l.Error(err)
			return nil, err
		}
	}

	return db, nil
}

func DeferClose(db *sqlx.DB, l *slog.SugaredLogger) {
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			l.Error(err)
			return
		}
	}(db)
}
