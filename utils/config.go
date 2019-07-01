package utils

type Config struct {
	Environment string `env:"ENVIRONMENT"`
	Username    string `env:"SUSERNAME"`
	Password    string `env:"SPASSWORD"`
}
