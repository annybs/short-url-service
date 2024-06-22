package cli

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type AddCmd struct {
	Token string `short:"t" help:"Bearer token (SHORTY_TOKEN)" default:"${token}"`
	URL   string `short:"u" help:"Shorty URL (SHORTY_URL)" default:"${url}"`

	Path        string `arg:"" help:"Redirect path"`
	Destination string `arg:"" help:"Destination URL"`
}

func (c *AddCmd) Run(ctx *Context) error {
	url := strings.Join([]string{c.URL, c.Path}, "")
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(c.Destination))
	if err != nil {
		return err
	}
	if c.Token != "" {
		req.Header.Add("authorization", fmt.Sprintf("bearer %s", c.Token))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if len(body) > 0 {
			return errors.New(string(body))
		}
		return errors.New(res.Status)
	}

	return nil
}

func (c *AddCmd) Validate() error {
	if len(c.Path) < 1 || c.Path[0] != '/' {
		return errors.New("invalid <path>")
	}
	return nil
}
