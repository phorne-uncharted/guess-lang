package env

import (
	"sync"

	"github.com/caarlos0/env"
)

var (
	cfg  *Config
	once sync.Once
)

// Config represents the application configuration state loaded from env vars.
type Config struct {
	AppPort            string  `env:"PORT" envDefault:"8090"`
	DatabaseURL        string  `env:"DATABASE_URL" envDefault:""`
	PostgresBatchSize  int     `env:"PG_BATCH_SIZE" envDefault:"1000"`
	PostgresDatabase   string  `env:"PG_DATABASE" envDefault:"guess"`
	PostgresHost       string  `env:"PG_HOST" envDefault:"localhost"`
	PostgresLogLevel   string  `env:"PG_LOG_LEVEL" envDefault:"none"`
	PostgresPassword   string  `env:"PG_PASSWORD" envDefault:"guess"`
	PostgresPort       int     `env:"PG_PORT" envDefault:"5432"`
	PostgresRandomSeed float64 `env:"PG_RANDOM_SEED" envDefault:"0.2"`
	PostgresUser       string  `env:"PG_USER" envDefault:"guess"`
}

// LoadConfig loads the config from the environment if necessary and returns a copy.
func LoadConfig() (Config, error) {
	var err error
	once.Do(func() {
		cfg = &Config{}
		err = env.Parse(cfg)
		if err != nil {
			cfg = &Config{}
		}
	})
	return *cfg, err
}
