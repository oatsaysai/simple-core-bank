package logger

import "errors"

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]any

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The system shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	InstanceZapLogger int = iota
	InstanceLogrusLogger
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance")
)

// Logger is our contract for the logger
type Logger interface {
	// Debugf(format string, args ...any)

	// Infof(format string, args ...any)

	// Warnf(format string, args ...any)

	// Errorf(format string, args ...any)

	// Fatalf(format string, args ...any)

	// Panicf(format string, args ...any)

	WithFields(keyValues Fields) Logger

	Fatal(v ...any)

	Fatalf(format string, v ...any)

	Fatalln(v ...any)

	Panic(v ...any)

	Panicf(format string, v ...any)

	Panicln(v ...any)

	Print(v ...any)

	Printf(format string, v ...any)

	Println(v ...any)

	Debug(args ...any)

	Debugf(format string, args ...any)

	Debugln(args ...any)

	Info(args ...any)

	Infof(format string, args ...any)

	Infoln(args ...any)

	Warn(args ...any)

	Warnf(format string, args ...any)

	Warnln(args ...any)

	Error(args ...any)

	Errorf(format string, args ...any)

	Errorln(args ...any)
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
	Color             bool
}

// NewLogger returns an instance of logger
func NewLogger(config *Configuration, loggerInstance int) (Logger, error) {
	if config == nil {
		config = &Configuration{
			EnableConsole:     true,
			ConsoleLevel:      "debug",
			ConsoleJSONFormat: false,
			EnableFile:        false,
			// FileLevel:         log.Info,
			// FileJSONFormat:    true,
			// FileLocation:      "log.log",
			Color: true,
		}
	}

	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := newZapLogger(*config)
		if err != nil {
			return nil, err
		}
		return logger, nil
	// case InstanceLogrusLogger:
	// 	logger, err := newLogrusLogger(*config)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return logger, nil

	default:
		return nil, errInvalidLoggerInstance
	}
}

func NormalizeLogLevel(logLevel string) string {
	var nomalizedLogLevel string
	switch logLevel {
	case "info":
		nomalizedLogLevel = Info
	case "debug":
		nomalizedLogLevel = Debug
	case "warn":
		nomalizedLogLevel = Warn
	case "error":
		nomalizedLogLevel = Error
	case "fatal":
		nomalizedLogLevel = Fatal
	default:
		nomalizedLogLevel = Debug
	}
	return nomalizedLogLevel
}
