package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port     int
	Database string
}

func loadConfig() *Config {
	cfg := &Config{
		Port:     1234,
		Database: "sqlite://myfile.db",
	}

	viper.SetConfigName("front-test")
	viper.SetEnvPrefix("ft")

	// Defaults
	viper.SetDefault("Port", 897)
	viper.SetDefault("Database", "")

	//Flags
	viper.AutomaticEnv()

	viper.ReadInConfig()

	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Println("cannot unmarshal config: %s", err)
	}

	return cfg
}

func main() {
	config := loadConfig()
	fmt.Println(config)
}
