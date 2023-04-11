package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/smartpcr/azs-2-tf/utils"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

type LoggerType int

const (
	ConsoleLogger LoggerType = iota
	FileLogger
)

var (
	Log *LoggerWrapper
	_   Logger = &LoggerWrapper{}
)

type Logger interface {
	Debug(msg string, fields ...KeyValuePair)
	Info(msg string, fields ...KeyValuePair)
	Warn(msg string, fields ...KeyValuePair)
	Error(msg string, fields ...KeyValuePair)
	Fatal(msg string, fields ...KeyValuePair)

	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
}

type LoggerWrapper struct {
	ctx        context.Context
	logger     *logrus.Logger
	loggerType LoggerType
}

type KeyValuePair struct {
	Key   string
	Value interface{}
}

func init() {
	appSettings := &utils.AppSettings{}
	Log = New(context.Background(), appSettings, ConsoleLogger, FileLogger)
}

func New(ctx context.Context, appSettings utils.Settings, loggerTypes ...LoggerType) *LoggerWrapper {
	logFolder := appSettings.GetLogFolderPath()
	err := utils.EnsureDirectory(logFolder)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation:    true,
		ForceColors:               true,
		PadLevelText:              true,
		FullTimestamp:             true,
		EnvironmentOverrideColors: true,
	})

	allWriters := make([]io.Writer, 0, len(loggerTypes))
	for _, loggerType := range loggerTypes {
		if loggerType == FileLogger {
			var fileLogger = &lumberjack.Logger{
				Filename:   filepath.Join(logFolder, appSettings.GetLogFileName()),
				MaxSize:    50, // megabytes
				MaxAge:     28, //days
				MaxBackups: 3,
				LocalTime:  true,
				Compress:   false,
			}
			allWriters = append(allWriters, fileLogger)
		} else if loggerType == ConsoleLogger {
			allWriters = append(allWriters, os.Stdout)
		}
	}
	logger.SetOutput(io.MultiWriter(allWriters...))

	return &LoggerWrapper{
		ctx:    ctx,
		logger: logger,
	}
}

func (l *LoggerWrapper) SetLevel(level logrus.Level) {
	l.logger.SetLevel(level)
}

func (l *LoggerWrapper) Debug(msg string, fields ...KeyValuePair) {
	l.addContextCommonFields(fields...)
	l.logger.WithFields(createMap(fields...)).Debug(msg)
}

func (l *LoggerWrapper) Info(msg string, fields ...KeyValuePair) {
	l.addContextCommonFields(fields...)
	l.logger.WithFields(createMap(fields...)).Info(msg)
}

func (l *LoggerWrapper) Warn(msg string, fields ...KeyValuePair) {
	l.addContextCommonFields(fields...)
	l.logger.WithFields(createMap(fields...)).Warn(msg)
}

func (l *LoggerWrapper) Error(msg string, fields ...KeyValuePair) {
	l.addContextCommonFields(fields...)
	l.logger.WithFields(createMap(fields...)).Error(msg)
}

func (l *LoggerWrapper) Fatal(msg string, fields ...KeyValuePair) {
	l.addContextCommonFields(fields...)
	l.logger.WithFields(createMap(fields...)).Fatal(msg)
}

func (l *LoggerWrapper) Debugf(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}

func (l *LoggerWrapper) Infof(msg string, args ...interface{}) {
	l.logger.Infof(msg, args...)
}

func (l *LoggerWrapper) Warnf(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args...)
}

func (l *LoggerWrapper) Errorf(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args...)
}

func (l *LoggerWrapper) Fatalf(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args...)
}

func (l *LoggerWrapper) addContextCommonFields(fields ...KeyValuePair) {
	mappedFields := createMap(fields...)
	if l.ctx != nil && l.ctx.Value("commonFields") != nil {
		for k, v := range l.ctx.Value("commonFields").(map[string]interface{}) {
			if _, ok := mappedFields[k]; !ok {
				mappedFields[k] = v
			}
		}
	}
}

func createMap(fields ...KeyValuePair) map[string]interface{} {
	m := make(map[string]interface{})
	if fields == nil {
		return m
	}

	for _, f := range fields {
		m[f.Key] = f.Value
	}
	return m
}
