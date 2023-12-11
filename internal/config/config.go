package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env  string `yaml:"env" env-default:"local"`
	Host string `yaml:"host" env-default:"localhost"`
	Port int    `yaml:"port" env-required:"true"`
}

// mustLoad - this fn must Load config file(required)
// If not load he'll PANIC
func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Error")
	}

	// TODO: Check file is exists or not
	var cfg Config

	configPath := fetchFlagOrEnv()

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("can't read config.yaml file" + err.Error())
	}

	return &cfg
}

func fetchFlagOrEnv() string {
	var configPath string

	flag.StringVar(&configPath, "config-path", "", "config-path")
	flag.Parse()

	if configPath != "" {
		return configPath
	}

	configPath = os.Getenv("CONFIG_PATH")
	if configPath != "" {
		return configPath
	}

	panic("ConfigPath is empty")
}
