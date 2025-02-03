package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger    *zap.Logger
	lumLogger *lumberjack.Logger

	level    string
	savePath string
	fileName string
	fileExt  string
}

// OptionFunc defines the function to change the default configuration
type OptionFunc func(*Logger)

// SetLogLevel to change the default level
func SetLogLevel(level string) OptionFunc {
	return func(l *Logger) {
		l.level = level
	}
}

// SetLogSavePath to change the default savePath
func SetLogSavePath(savePath string) OptionFunc {
	return func(l *Logger) {
		l.savePath = savePath
	}
}

// SetLogFileName to change the default fileName
func SetLogFileName(fileName string) OptionFunc {
	return func(l *Logger) {
		l.fileName = fileName
	}
}

// SetLogFileExt to change the default fileExt
func SetLogFileExt(fileExt string) OptionFunc {
	return func(l *Logger) {
		l.fileExt = fileExt
	}
}

// New returns a new blank logger instance
// By default, the configuration is:
// - level: "debug"
// - level: "./logs"
// - level: "log"
// - level: "log"
func New(opts ...OptionFunc) *Logger {
	l := &Logger{
		level:    "debug",
		savePath: "./logs",
		fileName: "log",
		fileExt:  "log",
	}

	return l.With(opts...).run()
}

// With returns a new logger instance with provided options
func (l *Logger) With(opts ...OptionFunc) *Logger {
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// run will initialize the lumlogger and zap
func (l *Logger) run() *Logger {
	l.newLumLogger()
	l.newZapLogger()
	return l
}

// new lumlogger
// set log rolling files
func (l *Logger) newLumLogger() {
	fileName := filepath.FromSlash(fmt.Sprintf("%s/%s.%s", l.savePath, l.fileName, l.fileExt))
	l.lumLogger = &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1,     // megabytes
		MaxBackups: 3,     // number of files backed up
		MaxAge:     28,    // number of days to keep files
		Compress:   false, // disabled by default
	}
}

// new zap logger
func (l *Logger) newZapLogger() {
	level := zap.NewAtomicLevelAt(l.getLevellogger(l.level))
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// set file and stdout encoder
	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	// tee core
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(l.lumLogger), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	l.logger = zap.New(
		core,
		zap.AddCaller(), // file line number
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	// switch to global
	// zap.L() or zap.S()
	zap.ReplaceGlobals(l.logger)
}

// get zap's log level from external settings
func (l *Logger) getLevellogger(level string) zapcore.Level {
	if level == "debug" {
		return zap.DebugLevel
	}

	return zap.InfoLevel
}

// close lumlogger & zap
func (l *Logger) Close() error {
	if err := l.closeLumLogger(); err != nil {
		return err
	}
	return l.syncZap()
}

// call lumlogger closes the current logfile
func (l *Logger) closeLumLogger() error {
	if l.lumLogger != nil {
		if err := l.lumLogger.Close(); err != nil {
			return err
		}
	}
	return nil
}

// call zap Sync before exiting
func (l *Logger) syncZap() error {
	if l.logger != nil {
		if err := l.logger.Sync(); err != nil {
			return err
		}
	}
	return nil
}
