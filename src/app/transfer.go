package app

import (
	"github.com/shopspring/decimal"
	log "repo.blockfint.com/sakkarin/go-http-server-template/src/logger"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/model"
)

func (ctx *Context) TransferIn(params *model.TransferInParams) (*model.TransferInResponse, error) {
	logger := ctx.Logger
	logger = logger.WithFields(log.Fields{
		"func": "TransferIn",
	})
	logger.Info("Begin")
	logger.Debugf("params: %+v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	txID, err := ctx.DB.TransferIn(
		ctx.FiberCtx.Context(),
		params.ToAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, err
	}

	return &model.TransferInResponse{
		TransactionID: *txID,
		ToAccountNo:   params.ToAccountNo,
		Amount:        params.Amount,
	}, nil
}

func (ctx *Context) TransferOut(params *model.TransferOutParams) (*model.TransferOutResponse, error) {
	logger := ctx.Logger
	logger = logger.WithFields(log.Fields{
		"func": "TransferOut",
	})
	logger.Info("Begin")
	logger.Debugf("params: %+v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	txID, err := ctx.DB.TransferOut(
		ctx.FiberCtx.Context(),
		params.FromAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, err
	}

	return &model.TransferOutResponse{
		TransactionID: *txID,
		FromAccountNo: params.FromAccountNo,
		Amount:        params.Amount,
	}, nil
}

func (ctx *Context) Transfer(params *model.TransferParams) (*model.TransferResponse, error) {
	logger := ctx.Logger
	logger = logger.WithFields(log.Fields{
		"func": "Transfer",
	})
	logger.Info("Begin")
	logger.Debugf("params: %+v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	txID, err := ctx.DB.Transfer(
		ctx.FiberCtx.Context(),
		params.FromAccountNo,
		params.ToAccountNo,
		decimal.NewFromFloat(params.Amount),
	)
	if err != nil {
		return nil, err
	}

	return &model.TransferResponse{
		TransactionID: *txID,
		FromAccountNo: params.FromAccountNo,
		ToAccountNo:   params.ToAccountNo,
		Amount:        params.Amount,
	}, nil
}
