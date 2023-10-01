package db

import (
	"best_friends_bot/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Repo struct {
	Message
	DB *pgxpool.Pool
}

func NewRepo(cfg *config.DBConfig) (*Repo, error) {
	db, err := CreateDatabasePoolConnections(cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize DB: %s", err)
	}

	return &Repo{
		Message: NewMessage(db),
		DB:      db,
	}, nil
}

func CreateDatabasePoolConnections(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(cfg.DBString)
	if err != nil {
		return nil, errors.Wrap(err, "err parse config for poolCfg")
	}
	poolCfg.ConnConfig.PreferSimpleProtocol = true //disable prepared queries for pgbouncer

	pool, err := pgxpool.ConnectConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "err init connection for new pool")
	}
	return pool, nil
}

func (m *Repo) Close() {
	m.DB.Close()
}
