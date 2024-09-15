package util

import (
	"flag"
	"os"

	"github.com/shubhamku044/management/model"
	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {
	Logger = *logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.SetOutput(os.Stdout)
}

func SetLogger() {
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level (debug, error, info, warning)")
	flag.Parse()
	switch *logLevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	case model.LogLevelWarn:
		Logger.SetLevel(logrus.WarnLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

func Log(logLevel, packageLevel, functionName string, message, paramerer interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if paramerer != nil {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v, paramerer: %v\n", packageLevel, functionName, message, paramerer)
		} else {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	case model.LogLevelError:
		if paramerer != nil {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v, paramerer: %v\n", packageLevel, functionName, message, paramerer)
		} else {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	case model.LogLevelWarn:
		if paramerer != nil {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v, paramerer: %v\n", packageLevel, functionName, message, paramerer)
		} else {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	default:
		if paramerer != nil {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v, paramerer: %v\n", packageLevel, functionName, message, paramerer)
		} else {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	}
}
