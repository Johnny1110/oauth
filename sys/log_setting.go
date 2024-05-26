package sys

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var sysLog *logrus.Logger

func init() {
	sysLog = logrus.New()
	sysLog.SetLevel(logrus.DebugLevel)

	sysLog.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   12,
	})

	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// 同时输出到文件和控制台
		sysLog.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		sysLog.Info("Failed to log to file, using default stderr")
	}

	sysLog.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func Logger() *logrus.Logger {
	return sysLog
}
