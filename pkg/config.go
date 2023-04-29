package pkg

import (
	"context"
	"sync"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	"github.com/sethvargo/go-envconfig"
	log "github.com/sirupsen/logrus"
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
	DBDatabase string `env:"DB_DATABASE"`
	DBHost     string `env:"DB_HOST"`
	DBLogLevel int    `env:"DB_LOG_LEVEL"`
	DBPassword string `env:"DB_PASSWORD"`
	DBPort     int    `env:"DB_PORT"`
	DBUsername string `env:"DB_USERNAME"`
}

// Config returns a new Conf struct with the configs
// Is a singleton with one memory address
func Config() *Conf {
	confOnce.Do(func() {
		conf = &Conf{}
		if err := envconfig.Process(context.Background(), conf); err != nil {
			log.Fatal(err)
		}
		log.Info(pp.Sprint(conf))
	})
	return conf
}

func dotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
