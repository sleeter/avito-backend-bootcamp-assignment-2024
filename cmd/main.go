package main

import (
	http_server "backend-bootcamp-assignment-2024/internal/http-server"
	"backend-bootcamp-assignment-2024/internal/pkg/pgdb"
	"backend-bootcamp-assignment-2024/internal/repository"
	"backend-bootcamp-assignment-2024/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"

	"backend-bootcamp-assignment-2024/internal/core"
	"backend-bootcamp-assignment-2024/internal/pkg/config"

	"github.com/avast/retry-go/v4"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	ctx := context.Background()

	//TODO: logger

	//TODO: config
	loader := config.PrepareLoader(config.WithConfigPath("./config/config.yaml"))
	cfg, err := core.ParseConfig(loader)
	if err != nil {

	}

	//TODO: repo, migrations
	err = retry.Do(func() error {
		return UpMigrations(cfg)
	}, retry.Attempts(4), retry.Delay(2*time.Second))
	if err != nil {

	}

	pool, err := pgxpool.Connect(ctx, cfg.Storage.URL)
	if err != nil {

	}
	qm := pgdb.NewQueryManager(pool)
	tm := pgdb.NewTransactionManager(pool)

	hRepo := repository.NewHouseRepository(qm)
	fRepo := repository.NewFlatRepository(qm)
	uRepo := repository.NewUserRepository(qm)

	//TODO: service
	hService := service.NewHouseService(hRepo, tm)
	fService := service.NewFlatService(fRepo, hRepo, tm)
	uService := service.NewUserService(uRepo, tm)
	//TODO: server
	app := http_server.New(&service.Service{
		HouseService: hService,
		FlatService:  fService,
		UserService:  uService,
	}, cfg)

	if err := app.Start(ctx); err != nil {

	}

	//TODO: graceful shutdown
}

func UpMigrations(cfg *core.Config) error {
	db, err := sql.Open("pgx", cfg.Storage.URL)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
