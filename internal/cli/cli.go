package cli

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Context struct {
	Config *viper.Viper
	Log    zerolog.Logger
}

type CLI struct {
	Add   AddCmd   `cmd:"" help:"Add a redirect"`
	Get   GetCmd   `cmd:"" help:"Get a redirect"`
	Rm    RmCmd    `cmd:"" help:"Delete a redirect"`
	Start StartCmd `cmd:"" help:"Start Shorty"`
}
