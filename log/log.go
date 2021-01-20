/**
  @author: xinyulu
  @date: 2021/1/19 13:57
  @note: 
**/
package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "io"
    "os"
)

var (
    logger      *zap.Logger
    atomicLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    writeSyncer = zapcore.AddSync(os.Stdout)
    encoderConfig zapcore.EncoderConfig
    encoder zapcore.Encoder
)

var levelMap = map[string]zapcore.Level{
    "debug":  zapcore.DebugLevel,
    "info":   zapcore.InfoLevel,
    "warn":   zapcore.WarnLevel,
    "error":  zapcore.ErrorLevel,
    "dpanic": zapcore.DPanicLevel,
    "panic":  zapcore.PanicLevel,
    "fatal":  zapcore.FatalLevel,
}

func init() {
    encoderConfig = getEncoderConfig()
    encoder = zapcore.NewJSONEncoder(encoderConfig)
    initLogger()
}

func getEncoderConfig() zapcore.EncoderConfig {
    return zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder,
        EncodeName:     zapcore.FullNameEncoder,
    }
}



func SetLevel(s string) {
    if level, ok := levelMap[s]; ok {
        atomicLevel.SetLevel(level)
        return
    }
    atomicLevel.SetLevel(zapcore.InfoLevel)
}

func initLogger() {
    core := zapcore.NewCore(encoder, writeSyncer, atomicLevel)
    logger = zap.New(core, zap.AddCaller(), zap.Development(), zap.AddCallerSkip(1))
}

func SetOutput(w io.Writer) {
    if w != nil {
        writeSyncer = zapcore.AddSync(w)
        initLogger()
        return
    }
}

func Field(key string, val interface{}) zap.Field {
    return zap.Any(key, val)
}

func Debug(msg string, fields ...zap.Field) {
    logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
    logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
    logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
    logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
    logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
    logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
    logger.Fatal(msg, fields...)
}
