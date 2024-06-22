package cli

import (
	"net"
	"net/http"

	"github.com/annybs/ezdb"
	"github.com/recipeer/short-url-service/internal/api"
)

type StartCmd struct {
	DatabasePath string `short:"d" help:"Database path (SHORTY_DATABASE_PATH)" default:"${database_path}"`
	Host         string `short:"h" help:"HTTP bind host (SHORTY_HOST)" default:"${host}"`
	Port         string `short:"p" help:"HTTP port (SHORTY_PORT)" default:"${port}"`
	Token        string `short:"t" help:"Bearer token (SHORTY_TOKEN)" default:"${token}"`
}

func (c *StartCmd) Run(ctx *Context) error {
	errc := make(chan error)

	db := ezdb.Memory(
		ezdb.LevelDB(c.DatabasePath, ezdb.Bytes(), nil),
	)
	if err := db.Open(); err != nil {
		return err
	}

	a := &api.API{
		DB:  db,
		Log: ctx.Log,

		Token: c.Token,
	}

	addr := net.JoinHostPort(c.Host, c.Port)
	go func() {
		errc <- http.ListenAndServe(addr, a)
	}()

	ctx.Log.Info().Msgf("Listening at %s", addr)
	return <-errc
}
