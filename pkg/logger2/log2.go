package logger2

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func CustomEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[level])
}

// zap ensure that zap.Logger is safe for concurrent use
type WLogger struct {
	l     *zap.Logger
}

type Level = zapcore.Level
type Field = zap.Field

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

func (l *WLogger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *WLogger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *WLogger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *WLogger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}
func (l *WLogger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}
func (l *WLogger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}
func (l *WLogger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString

	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Durationp   = zap.Durationp

	Any         = zap.Any

	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug
)

var std = New(os.Stderr, InfoLevel)

func New(writer io.Writer, level Level) *WLogger{
	//Define config for the console output
	cfgConsole := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "severity",
		EncodeLevel:  CustomEncodeLevel,
		TimeKey:      "time",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		CallerKey:    "caller",
		//EncodeCaller: zapcore.FullCallerEncoder, //For full path
		EncodeCaller: zapcore.ShortCallerEncoder,  //For short path
		FunctionKey: "func",          			   ////Print function name
	}

	consoleDebugging := zapcore.Lock(os.Stdout)
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(cfgConsole), consoleDebugging, zap.DebugLevel))

	wlogger := &WLogger{
		l:     zap.New(core,zap.AddCaller(),zap.AddCallerSkip(1)),
	}
	return wlogger
}

func (l *WLogger) Sync() error {
	return l.l.Sync()
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}