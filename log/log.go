package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func Init() {
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetReportCaller(true)
}
