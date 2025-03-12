package postgres

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/tracelog"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/constant"
	log "repo.blockfint.com/sakkarin/go-http-server-template/src/logger"
)

type PostgresLogger struct {
	Logger log.Logger
}

func NewDatabaseLogger(logger *log.Logger) *PostgresLogger {
	return &PostgresLogger{Logger: *logger}
}

func (pglog *PostgresLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {

	// idea from https://github.com/jackc/pgx-logrus/blob/master/adapter.go

	var logger = pglog.Logger

	// Add span_id to log
	if ctx.Value(constant.ContextKeySpanID) != nil {
		logger = logger.WithFields(log.Fields{
			"span_id": ctx.Value(constant.ContextKeySpanID),
		})
	}

	// Add trace_id to log
	if ctx.Value(constant.ContextKeyTraceID) != nil {
		logger = logger.WithFields(log.Fields{
			"trace_id": ctx.Value(constant.ContextKeyTraceID),
		})
	}

	// Add user_id to log
	if ctx.Value(constant.ContextKeyUserID) != nil {
		logger = logger.WithFields(log.Fields{
			"user_id": ctx.Value(constant.ContextKeyUserID),
		})
	}

	// Add username to log
	if ctx.Value(constant.ContextKeyUsername) != nil {
		logger = logger.WithFields(log.Fields{
			"username": ctx.Value(constant.ContextKeyUsername),
		})
	}

	if data != nil {
		if data["sql"] != nil {
			sqlTemp := data["sql"].(string)
			sqlTemp = strings.ReplaceAll(sqlTemp, "\n\t\t", " ")
			sqlTemp = strings.ReplaceAll(sqlTemp, "\t\t\t", " ")
			sqlTemp = strings.ReplaceAll(sqlTemp, "\t", "")
			sqlTemp = strings.ReplaceAll(sqlTemp, "\n", "")
			data["sql"] = sqlTemp
		}
		logger = logger.WithFields(data)
	}

	switch level {
	case tracelog.LogLevelTrace:
		logger.WithFields(createFields("PGX_LOG_LEVEL", level)).Debugf(msg)
	case tracelog.LogLevelDebug:
		logger.Debugf(msg)
	case tracelog.LogLevelInfo:
		// Log info with debug level
		logger.Debugf(msg)
	case tracelog.LogLevelWarn:
		logger.Warnf(msg)
	case tracelog.LogLevelError:
		logger.Errorf(msg)
	default:
		logger.WithFields(createFields("INVALID_PGX_LOG_LEVEL", level)).Errorf(msg)
	}
}

func createFields(key string, value any) log.Fields {
	var fieldMap = make(map[string]any)
	fieldMap[key] = value
	return fieldMap
}
