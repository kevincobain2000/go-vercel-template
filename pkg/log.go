package pkg

import (
	"github.com/charmbracelet/log"
)

func init() {
	log.Info("log init")

	log.SetLevel(Config().LogLevel)
	log.SetReportCaller(true)

	log.Info("log", "level", Config().LogLevel)
	log.Info("log", "set reporter", true)
}
