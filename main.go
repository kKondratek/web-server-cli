package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {

	const version = "0.1.0"

	app := &cli.App{
		Name:    "web-server-cli",
		Usage:   "Simple CLI application to run local web server",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run local http server",
				Action: func(c *cli.Context) error {
					var filepath string = c.String("file")
					filepath, err := validatePath(filepath)

					if err != nil {
						log.Fatal("index.html not found in given path")
					}

					fmt.Println("Starting server on http://localhost:4000")
					http.Handle("/", http.FileServer(http.Dir(filepath)))
					http.ListenAndServe(":4000", nil)

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "specify location of index.html file",
						Required: true,
					},
					// TODO: --port flag
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Print the version",
				Action: func(c *cli.Context) error {
					fmt.Println("web-server-cli version " + version)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// checks if index.html exists and returns its path
func validatePath(path string) (string, error) {
	var resultPath string

	// checks if provided path contains index.html
	pathWithFile, _ := regexp.Compile("(.*)index.html")
	if pathWithFile.MatchString(path) {
		resultPath = strings.TrimSuffix(path, "index.html")
	} else {
		resultPath = path
		if isWindows() {
			path += "\\index.html"
		} else {
			path += "/index.html"
		}
	}
	// checks if index.html exsists in given path
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}

	return resultPath, nil
}

// checks if program is running on Windows os
func isWindows() bool {
	return runtime.GOOS == "windows"
}
