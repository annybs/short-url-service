package cli

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GetCmd struct {
	URL string `short:"u" help:"Shorty URL (SHORTY_URL)" default:"${url}"`

	Path string `arg:"" help:"Redirect path"`
}

func (c *GetCmd) Run(ctx *Context) error {
	url := strings.Join([]string{c.URL, c.Path}, "")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// Create a specialized client that ignores redirection
	client := &http.Client{
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusNotFound {
		return nil
	} else if res.StatusCode != http.StatusPermanentRedirect {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if len(body) > 0 {
			return errors.New(string(body))
		}
		return errors.New(res.Status)
	}

	dest := res.Header.Get("location")
	fmt.Println(dest)

	return nil
}

func (c *GetCmd) Validate() error {
	if len(c.Path) < 1 || c.Path[0] != '/' {
		return errors.New("invalid <path>")
	}
	return nil
}
