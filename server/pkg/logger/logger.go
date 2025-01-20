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
}

func NewLogger() *Logger {
	l := &Logger{}
	l.newLumLogger()
	l.newZapLogger()

	return l
}

// new lumlogger
// set log rolling files
func (l *Logger) newLumLogger() {
	fileName := filepath.FromSlash(fmt.Sprintf("%s/%s.%s", os.Getenv("LOG_SAVE_PATH"), os.Getenv("LOG_FILE_NAME"), os.Getenv("LOG_FILE_EXT")))
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
	level := zap.NewAtomicLevelAt(l.getLevellogger(os.Getenv("LOG_LEVEL")))
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
