package postgres

import (
	"context"

	"repo.blockfint.com/sakkarin/go-http-server-template/src/model"
)

func (pgdb *PostgresqlDB) GetTransactionByAccountNo(ctx context.Context, accountNo string) ([]model.Transaction, error) {

	txs := []model.Transaction{}
	err := pgdb.DB.QueryRow(
		ctx,
		`
		SELECT
			COALESCE(jsonb_agg(d.*), '[]') as rows
		FROM
			(
				SELECT *
				FROM transactions
				WHERE from_account_no = $1
					OR to_account_no = $1
				ORDER BY id
			) as d
		`,
		accountNo,
	).Scan(&txs)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
