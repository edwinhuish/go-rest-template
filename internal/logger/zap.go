package logger

import (
	"fmt"

	"github.com/edwinhuish/go-rest-template/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	//"github.com/natefinch/lumberjack"
	"os"
	"path"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerInterface interface {
	Debugt(msg string, fields ...zapcore.Field)
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Debugs(args ...interface{})

	Infot(msg string, fields ...zapcore.Field)
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Infos(args ...interface{})

	Warnt(msg string, fields ...zapcore.Field)
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Warns(args ...interface{})

	Errort(msg string, fields ...zapcore.Field)
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Errors(args ...interface{})

	Panict(msg string, fields ...zapcore.Field)
	Panicf(template string, args ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Panic(msg string, keysAndValues ...interface{})
	Panics(args ...interface{})

	Fatalt(msg string, fields ...zapcore.Field)
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
	Fatals(args ...interface{})

	AtLevel(level zapcore.Level, msg string, fields ...zapcore.Field) *Logger
}

type Logger struct {
	Unsugared *zap.Logger
	*zap.SugaredLogger
}

func (logger *Logger) AtLevel(level zapcore.Level, msg string, fields ...zapcore.Field) *Logger {
	switch level {
	case zapcore.DebugLevel:
		logger.Unsugared.Debug(msg, fields...)
	case zapcore.PanicLevel:
		logger.Unsugared.Panic(msg, fields...)
	case zapcore.ErrorLevel:
		logger.Unsugared.Error(msg, fields...)
	case zapcore.WarnLevel:
		logger.Unsugared.Warn(msg, fields...)
	case zapcore.InfoLevel:
		logger.Unsugared.Info(msg, fields...)
	case zapcore.FatalLevel:
		logger.Unsugared.Fatal(msg, fields...)
	default:
		logger.Unsugared.Warn("Logging at unknown level", zap.Any("level", level))
		logger.Unsugared.Warn(msg, fields...)
	}
	return logger
}

// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func (logger *Logger) Debugt(msg string, fields ...zapcore.Field) {
	logger.Unsugared.Debug(msg, fields...)
}

func (logger *Logger) Debug(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.Debugw(msg, keysAndValues...)
}

func (logger *Logger) Debugs(args ...interface{}) {
	logger.SugaredLogger.Debug(args...)
}

// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func (logger *Logger) Infot(msg string, fields ...zapcore.Field) {
	logger.Unsugared.Info(msg, fields...)
}

func (logger *Logger) Info(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.Infow(msg, keysAndValues...)
}

func (logger *Logger) Infos(args ...interface{}) {
	logger.SugaredLogger.Info(args...)
}

// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func (logger *Logger) Warnt(msg string, fields ...zapcore.Field) {
	logger.Unsugared.Warn(msg, fields...)
}

func (logger *Logger) Warn(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.Warnw(msg, keysAndValues...)
}

func (logger *Logger) Warns(args ...interface{}) {
	logger.SugaredLogger.Warn(args...)
}

// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func (logger *Logger) Errort(msg string, fields ...zapcore.Field) {
	logger.Unsugared.Error(msg, fields...)
}

func (logger *Logger) Error(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.Errorw(msg, keysAndValues...)
}

func (logger *Logger) Errors(args ...interface{}) {
	logger.SugaredLogger.Error(args...)
}

// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func (logger *Logger) Panict(msg string, fields ...zapcore.Field) {
	logger.Unsugared.Panic(msg, fields...)
}

func (logger *Logger) Panic(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.Panicw(msg, keysAndValues...)
}

func (logger *Logger) Panics(args ...interface{}) {
	logger.SugaredLogger.Panic(args...)
}

// Use zap.String(key, value), zap.Int(key, value) to log fields. These fields
// will be marshalled as JSON in the logfile and key value pairs in the console!
func (logger *Logger) Fatalt(msg string, fields ...zapcore.Field) {
	logger.Unsugared.Fatal(msg, fields...)
}

func (logger *Logger) Fatal(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.Fatalw(msg, keysAndValues...)
}

func (logger *Logger) Fatals(args ...interface{}) {
	logger.SugaredLogger.Fatal(args...)
}

// How to log, by example:
// logger.Infot("Importing new file", zap.String("source", filename), zap.Int("size", 1024))
// logger.Info("Importing new file", "source", filename, "size", 1024)
// To log a stacktrace:
// logger.Errort("It went wrong, zap.Stack())

// defaultZapLogger is the default logger instance that should be used to log
// It's assigned a default value here for tests (which do not call log.Configure())
var defaultZapLogger *Logger

func getDefaultLogger() *Logger {

	if defaultZapLogger == nil {
		conf := config.GetConfig().Logger
		defaultZapLogger = newZapLogger(conf)
	}

	return defaultZapLogger
}

func Debugt(msg string, fields ...zapcore.Field) {
	getDefaultLogger().Debugt(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	getDefaultLogger().Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Debugw(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Debug(msg, keysAndValues...)
}

func Debugs(args ...interface{}) {
	getDefaultLogger().Debugs(args...)
}

func Infot(msg string, fields ...zapcore.Field) {
	getDefaultLogger().Infot(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	getDefaultLogger().Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Infow(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Info(msg, keysAndValues...)
}

func Infos(args ...interface{}) {
	getDefaultLogger().Infos(args...)
}

func Warnt(msg string, fields ...zapcore.Field) {
	getDefaultLogger().Warnt(msg, fields...)
}

func Warnf(template string, args ...interface{}) {
	getDefaultLogger().Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Warnw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Warn(msg, keysAndValues...)
}

func Warns(args ...interface{}) {
	getDefaultLogger().Warns(args...)
}

func Errort(msg string, fields ...zapcore.Field) {
	getDefaultLogger().Errort(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	fmt.Printf("这是FMT的输出："+template, args...)
	getDefaultLogger().Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Errorw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Error(msg, keysAndValues...)
}

func Errors(args ...interface{}) {
	getDefaultLogger().Errors(args...)
}

func Panict(msg string, fields ...zapcore.Field) {
	getDefaultLogger().Panict(msg, fields...)
}

func Panicf(template string, args ...interface{}) {
	getDefaultLogger().Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Panicw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Panic(msg, keysAndValues...)
}

func Panics(args ...interface{}) {
	getDefaultLogger().Panics(args...)
}

func Fatalt(msg string, fields ...zapcore.Field) {
	getDefaultLogger().Fatalt(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	getDefaultLogger().Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Fatalw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	getDefaultLogger().Fatal(msg, keysAndValues...)
}

func Fatals(args ...interface{}) {
	getDefaultLogger().Fatals(args...)
}

// AtLevel logs the message at a specific log level
func AtLevel(level zapcore.Level, msg string, fields ...zapcore.Field) {
	switch level {
	case zapcore.DebugLevel:
		Debugt(msg, fields...)
	case zapcore.PanicLevel:
		Panict(msg, fields...)
	case zapcore.ErrorLevel:
		Errort(msg, fields...)
	case zapcore.WarnLevel:
		Warnt(msg, fields...)
	case zapcore.InfoLevel:
		Infot(msg, fields...)
	case zapcore.FatalLevel:
		Fatalt(msg, fields...)
	default:
		Warnt("Logging at unknown level", zap.Any("level", level))
		Warnt(msg, fields...)
	}
}

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to configuration set.
func Configure(config config.LoggerConfig) *Logger {
	logger := newZapLogger(config)
	logger.Infot("logging configured",
		zap.Bool("consoleEnabled", config.ConsoleEnabled),
		zap.String("consoleLevel", config.ConsoleLevel),
		zap.Bool("consoleJson", config.ConsoleJson),
		zap.Bool("fileEnabled", config.FileEnabled),
		zap.String("fileLevel", config.FileLevel),
		zap.Bool("fileJson", config.FileJson),
		zap.String("logDirectory", config.Directory),
		zap.String("fileName", config.Filename),
		zap.Int("maxSizeMB", config.MaxSize),
		zap.Int("maxBackups", config.MaxBackups),
		zap.Int("maxAgeInDays", config.MaxAge))

	defaultZapLogger = logger
	zap.RedirectStdLog(getDefaultLogger().Unsugared)

	return logger
}

func newRollingFile(config config.LoggerConfig) zapcore.WriteSyncer {
	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		Error("can't create log directory", zap.Error(err), zap.String("path", config.Directory))
		return nil
	}

	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
		MaxBackups: config.MaxBackups, // files
	})
}

func newZapLogger(config config.LoggerConfig) *Logger {

	fmt.Printf("newZapLogger: %v", config)

	var consoleLevel zapcore.Level
	consoleLevel.Set(strings.ToLower(config.ConsoleLevel))

	var fileLevel zapcore.Level
	fileLevel.Set(strings.ToLower(config.FileLevel))

	consoleEncCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}
	jsonEncCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}

	consoleLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= consoleLevel
	})
	fileLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= fileLevel
	})

	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncCfg)
	fileEncoder := zapcore.NewJSONEncoder(jsonEncCfg)

	var cores []zapcore.Core

	if config.ConsoleEnabled {
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stderr), consoleLevelEnabler))
	}
	if config.FileEnabled {
		cores = append(cores, zapcore.NewCore(fileEncoder, newRollingFile(config), fileLevelEnabler))
	}
	core := zapcore.NewTee(cores...)

	unsugared := zap.New(core)
	return &Logger{
		Unsugared:     unsugared,
		SugaredLogger: unsugared.Sugar(),
	}
}
