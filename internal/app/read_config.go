package app

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// config file path stored in environment variable
const envConfigFilePath = "CONFIG_PATH"

// ReadConfig reads config from file
func ReadConfig(cfg any) error {
	configPath := os.Getenv(envConfigFilePath) // config.yml path stored in environ

	v := viper.New()
	v.SetConfigFile(configPath)

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	if err := v.Unmarshal(cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
