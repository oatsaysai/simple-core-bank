package logger

import "errors"

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

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
	// Debugf(format string, args ...interface{})

	// Infof(format string, args ...interface{})

	// Warnf(format string, args ...interface{})

	// Errorf(format string, args ...interface{})

	// Fatalf(format string, args ...interface{})

	// Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger

	Fatal(v ...interface{})

	Fatalf(format string, v ...interface{})

	Fatalln(v ...interface{})

	Panic(v ...interface{})

	Panicf(format string, v ...interface{})

	Panicln(v ...interface{})

	Print(v ...interface{})

	Printf(format string, v ...interface{})

	Println(v ...interface{})

	Debug(args ...interface{})

	Debugf(format string, args ...interface{})

	Debugln(args ...interface{})

	Info(args ...interface{})

	Infof(format string, args ...interface{})

	Infoln(args ...interface{})

	Warn(args ...interface{})

	Warnf(format string, args ...interface{})

	Warnln(args ...interface{})

	Error(args ...interface{})

	Errorf(format string, args ...interface{})

	Errorln(args ...interface{})
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
