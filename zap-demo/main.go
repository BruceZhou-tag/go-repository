package main

import (
	"net/http"
	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	for i :=0;i<10000;i++ {
		logger.Info("test rotate log.....")
	}
	simpleHttpGet("www.5lmh.com")
	simpleHttpGet("http://www.sogou.com")
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// logger, _ = zap.NewProduction()
	logger = zap.New(core,zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	return zapcore.NewJSONEncoder(encoderConfig)
}

// func getLogWriter() zapcore.WriteSyncer {
// 	file, _ := os.Create("./test.log")
// 	return zapcore.AddSync(file)
// }
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename :"./test.log",
		MaxSize: 1,
		MaxBackups:5,
		MaxAge: 30,
		Compress: false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
