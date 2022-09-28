package logs

import (
	"fmt"
	"gofun/conf"
	"io"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// 初始化日志
func InitConfig() (*zap.Logger, error) {

	// warnlevel以下属于info
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})
	// warnlevel及以上属于warn
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoWriter := getWriter(conf.Config.Log.Dir + "/" + "info.log")
	errorWriter := getWriter(conf.Config.Log.Dir + "/" + "error.log")
	encode := getCommonEncoder()
	// 设置日志级别（默认info级别，可以根据需要设置级别）
	// level := getLevel(conf.Config.Log.Level)
	core := zapcore.NewTee(
		zapcore.NewCore(encode, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encode, zapcore.AddSync(errorWriter), warnLevel),
	)

	// 设置初始化字段
	// filed := zap.Fields(zap.String("clientIp", ""))

	// AddCallerSkip增加了调用者注释跳过的调用者数量 显示调用打印日志的是哪一行的code行数
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger, nil
}

// 获取日志级别
func getLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// 格式化日期
func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func getWriter(filename string) io.Writer {
	// filename = strings.Replace(filename, ".log", "_", -1) + time.Now().AddDate(0, 0, -1).Format("20060102") + ".log"
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,                   // 日志文件的位置
		MaxSize:    conf.Config.Log.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: conf.Config.Log.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     conf.Config.Log.MaxAge,     // 保留旧文件的最大天数
		Compress:   conf.Config.Log.Compress,   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getProductEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = formatEncodeTime
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func geDevelopEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = formatEncodeTime
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getCommonEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",     // 时间
		LevelKey:       "level",  // 输出日志级别的key名
		NameKey:        "logger", // 输出日志名称的key名
		CallerKey:      "caller", // 输出日志调用者的key名
		MessageKey:     "type",   // 输入信息的key名
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器, 将日志级别字符串转化为小写
		EncodeTime:     formatEncodeTime,               // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行消耗的时间转化成浮点型的秒
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器FullCallerEncoder, 以包/文件:行号 格式化调用堆栈ShortCallerEncoder
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func One() *zap.Logger {
	return logger
}

func Debug(msg string, args ...zap.Field) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...zap.Field) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...zap.Field) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...zap.Field) {
	logger.Error(msg, args...)
}

func Panic(msg string, args ...zap.Field) {
	logger.Panic(msg, args...)
}

func Fatal(msg string, args ...zap.Field) {
	logger.Fatal(msg, args...)
}
