package pkg

import (
	"os"
	"strconv"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func init() {
	log.Info("dot env init")
	dotEnv()

	log.Info("config init")
	Config()
}

var (
	conf     *Conf
	confOnce sync.Once
)

// Conf is the configuration struct that returns all the configs
type Conf struct {
	DBDatabase string
	DBHost     string
	DBLogLevel int
	DBPassword string
	DBPort     int
	DBUsername string
	LogLevel   log.Level
}

// Config returns a new Conf struct with the configs
// Is a singleton with one memory address
func Config() *Conf {
	confOnce.Do(func() {
		dbLogLevel, err := strconv.Atoi(os.Getenv("DB_LOG_LEVEL"))
		if err != nil {
			log.Fatal(err)
		}
		dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			log.Fatal(err)
		}
		conf = &Conf{
			DBDatabase: os.Getenv("DB_DATABASE"),
			DBHost:     os.Getenv("DB_HOST"),
			DBLogLevel: dbLogLevel,
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBPort:     dbPort,
			DBUsername: os.Getenv("DB_USERNAME"),
			LogLevel:   log.DebugLevel,
		}
	})
	return conf
}

func dotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
