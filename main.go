package main

import (
	"github.com/alecthomas/kong"
	"github.com/recipeer/short-url-service/internal/cli"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func main() {
	// Initialise config from environment
	config := viper.New()
	config.SetEnvPrefix("shorty")
	config.SetDefault("database_path", ".shorty")
	config.SetDefault("port", "3000")
	config.SetDefault("url", "http://localhost:3000")
	config.AutomaticEnv()

	// Create logger
	log := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()

	// Initialise and run CLI
	vars := kong.Vars{
		"database_path": config.GetString("database_path"),
		"host":          config.GetString("host"),
		"port":          config.GetString("port"),
		"token":         config.GetString("token"),
		"url":           config.GetString("url"),
	}
	ctx := kong.Parse(&cli.CLI{}, vars)
	err := ctx.Run(&cli.Context{Config: config, Log: log})
	ctx.FatalIfErrorf(err)
}
