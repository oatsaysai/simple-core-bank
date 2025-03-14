package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func getEncoder(isJSON bool, color bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	if color {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Debug:
		return zapcore.DebugLevel
	case Error:
		return zapcore.ErrorLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZapLogger(config Configuration) (Logger, error) {
	cores := []zapcore.Core{}

	if config.EnableConsole {
		level := getZapLevel(config.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.ConsoleJSONFormat, config.Color), writer, level)
		cores = append(cores, core)
	}

	if config.EnableFile {
		level := getZapLevel(config.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.FileLocation,
			MaxSize:  100,
			Compress: true,
			MaxAge:   28,
		})
		core := zapcore.NewCore(getEncoder(config.FileJSONFormat, config.Color), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(1),
		zap.AddCaller(),
	).Sugar()

	return &zapLogger{
		sugaredLogger: logger,
	}, nil
}

func (l *zapLogger) Debugf(format string, args ...any) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...any) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...any) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...any) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...any) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...any) {
	l.sugaredLogger.Panicf(format, args...)
}

func (l *zapLogger) Debug(args ...any) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Debugln(args ...any) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Error(args ...any) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Errorln(args ...any) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Fatal(args ...any) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Fatalln(args ...any) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Info(args ...any) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Warn(args ...any) {
	l.sugaredLogger.Warn(args...)
}

func (l *zapLogger) Warnln(args ...any) {
	l.sugaredLogger.Warn(args...)
}

func (l *zapLogger) Infoln(args ...any) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Panic(args ...any) {
	l.sugaredLogger.Panic(args...)
}

func (l *zapLogger) Panicln(args ...any) {
	l.sugaredLogger.Panic(args...)
}

func (l *zapLogger) Print(args ...any) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Println(args ...any) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Printf(format string, args ...any) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]any, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}
