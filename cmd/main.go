package main

import (
	http_server "backend-bootcamp-assignment-2024/internal/http-server"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"backend-bootcamp-assignment-2024/internal/pkg/pgdb"
	"backend-bootcamp-assignment-2024/internal/repository"
	"backend-bootcamp-assignment-2024/internal/service"
	"backend-bootcamp-assignment-2024/pkg/sender"
	"context"
	"database/sql"
	"errors"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
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

	loader := config.PrepareLoader(config.WithConfigPath("./config/config.yaml"))
	cfg, err := core.ParseConfig(loader)
	if err != nil {
		log.Fatal(err)
	}

	err = retry.Do(func() error {
		return UpMigrations(cfg)
	}, retry.Attempts(4), retry.Delay(2*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.Connect(ctx, cfg.Storage.URL)
	if err != nil {
		log.Fatal(err)
	}
	qm := pgdb.NewQueryManager(pool)
	tm := pgdb.NewTransactionManager(pool)

	cacheForFlat := expirable.NewLRU[int32, []entity.Flat](0, nil, time.Duration(60)*time.Second)

	hRepo := repository.NewHouseRepository(qm)
	fRepo := repository.NewFlatRepository(qm, cacheForFlat)
	uRepo := repository.NewUserRepository(qm)
	sRepo := repository.NewSubscriberRepository(qm)

	s := sender.New()
	hService := service.NewHouseService(hRepo, tm)
	sService := service.NewSubscriberService(sRepo, tm, s)
	fService := service.NewFlatService(fRepo, hService, sService, tm)
	uService := service.NewUserService(uRepo, tm)

	app := http_server.New(&service.Service{
		HouseService:      hService,
		FlatService:       fService,
		UserService:       uService,
		SubscriberService: sService,
	}, cfg)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
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
