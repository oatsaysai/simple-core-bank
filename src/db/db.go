package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
	postgresql "repo.blockfint.com/sakkarin/go-http-server-template/src/db/postgres"
	log "repo.blockfint.com/sakkarin/go-http-server-template/src/logger"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/model"
)

type DB interface {
	GetAccount(ctx context.Context, accountNo string) (*string, *string, *decimal.Decimal, error)
	AccountExists(ctx context.Context, accountNo string) (bool, error)
	GetAccountNoAndInsertAccount(ctx context.Context, accountName string, balance decimal.Decimal) (string, error)
	PreGenerateAccountNo(ctx context.Context, batchSize int) error
	GetTransactionByAccountNo(ctx context.Context, accountNo string) ([]model.Transaction, error)
	TransferIn(ctx context.Context, toAccountNo string, amount decimal.Decimal) (*int64, error)
	TransferOut(ctx context.Context, fromAccountNo string, amount decimal.Decimal) (*int64, error)
	Transfer(ctx context.Context, fromAccountNo, toAccountNo string, amount decimal.Decimal) (*int64, error)

	Close() error
}

type PostgresqlDB struct {
	logger log.Logger
	DB     *pgxpool.Pool
}

func New(config *Config, logger log.Logger) (db DB, err error) {
	switch config.DBType {
	case "postgres":
		dbConfig, err := postgresql.InitConfig()
		if err != nil {
			return nil, err
		}

		return postgresql.New(dbConfig, logger)
	}

	return nil, errors.New("unsupported database type")
}

func (pgdb *PostgresqlDB) Close() error {
	pgdb.DB.Close()
	return nil
}
