package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/viper"
)

// ListenHTTP starts an HTTP server that redirects any request for a recognised path to the corresponding URL.
// The server address and port can be configured with HOST and PORT environment variables, respectively.
// The default port is 3000.
func ListenHTTP(config *viper.Viper, redirects map[string]string) <-chan error {
	errc := make(chan error)

	addr := net.JoinHostPort(config.GetString("host"), config.GetString("port"))
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			path := req.URL.Path
			dest := redirects[path]
			code := 404
			if dest != "" {
				code = 308
				w.Header().Set("Location", dest)
			}
			w.WriteHeader(code)
			w.Write([]byte{})
			fmt.Println(path, code, dest)
		})
		errc <- http.ListenAndServe(addr, http.DefaultServeMux)
	}()

	fmt.Println("Listening at", addr)
	return errc
}

// ReadRedirects parses a map of redirects from the CSV environment variable (literally, "CSV").
// The CSV data must be in the format:
//
//	path,url
//	/some/path,https://some-url.com
//
// The first line is always skipped, allowing for CSV headings.
func ReadRedirects(config *viper.Viper) (map[string]string, error) {
	data := config.GetString("csv")
	if data == "" {
		return nil, errors.New("no CSV data")
	}
	reader := csv.NewReader(strings.NewReader(data))
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	redirects := map[string]string{}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) != 2 {
			return nil, fmt.Errorf("invalid line %d", i)
		}
		redirects[row[0]] = row[1]
	}
	return redirects, nil
}

func main() {
	config := viper.New()
	config.SetConfigFile(".env")
	config.SetDefault("port", "3000")
	config.ReadInConfig()
	config.AutomaticEnv()

	redirects, err := ReadRedirects(config)
	if err != nil {
		panic(err)
	}

	errc := ListenHTTP(config, redirects)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errc:
		fmt.Println(err)
		return
	case <-sigc:
		fmt.Println("Stopped")
		return
	}
}
