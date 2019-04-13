package config

import (
	"github.com/caarlos0/env"
	"github.com/go-playground/log"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string `env:"APP_NAME" envDefault:"featherr"`
	AppEnv    string `env:"APP_ENV" envDefault:"local"`
	AppDebug  string `env:"APP_DEBUG" envDefault:"true"`
	AppPort   string `env:"APP_PORT" envDefault:"8000"`
	AppDomain string `env:"APP_DOMAIN" envDefault:"localhost"`

	DBName string `env:"DB_NAME" envDefault:"localhost"`
	DBPass string `env:"DB_PASS" envDefault:"test"`
	DBUser string `env:"DB_USER" envDefault:"root"`
	DBType string `env:"DB_TYPE" envDefault:"mysql"`
	DBHost string `env:"DB_HOST" envDefault:"localhost"`
	DBPort string `env:"DB_PORT" envDefault:"3306"`

	JWTSecret    string `env:"JWT_SECRET" envDefault:"test"`
	SpacesKey    string `env:"SPACES_KEY" envDefault:"test"`
	SpacesSecret string `env:"SPACES_SECRET" envDefault:"test"`
}

var cfg *Config

// Parse parses, validates and then returns the application
// configuration based on ENV variables
func init() {
	if err := godotenv.Load("/Users/abdullahimahamed/go/src/github.com/featherr-engineering/rest-api/.env"); err != nil {
		log.Warn("File .env not found, reading configuration from ENV")
	}

	log.Info("Parsing ENV vars...")
	defer log.Info("Done Parsing ENV vars")

	cfg = &Config{}

	if err := env.Parse(cfg); err != nil {
		log.WithFields(log.F("error", err)).Warn("Errors Parsing Configuration")
	}

	return
}

func GetConfig() *Config {
	return cfg
}
