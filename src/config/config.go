package config

type Config struct {
	Environment string `env:"APP_ENV,required"`
}
