package app

import (
	log "repo.blockfint.com/sakkarin/go-http-server-template/src/logger"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/model"
)

func (ctx *Context) GetTransactionByAccountNo(params *model.GetTransactionParams) ([]model.Transaction, error) {
	logger := ctx.Logger
	logger = logger.WithFields(log.Fields{
		"func": "GetTransactionByAccountNo",
	})
	logger.Info("Begin")
	logger.Debugf("params: %+v", params)
	defer logger.Info("End")

	if err := ValidateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	return ctx.DB.GetTransactionByAccountNo(ctx.FiberCtx.Context(), params.AccountNo)
}
