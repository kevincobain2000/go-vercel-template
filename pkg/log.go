package pkg

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("log init")

	for _, arg := range os.Args {
		if arg == "-v" {
			log.SetReportCaller(true)
			log.SetLevel(log.DebugLevel)
		}
	}

	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		PadLevelText:    true,
	})

	log.Info("log level " + log.DebugLevel.String())
	log.Info("log reporter true")
}
