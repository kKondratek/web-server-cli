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

	app := &cli.App{
		// TODO: main command
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run local http server",
				Action: func(c *cli.Context) error {
					var filepath string = c.String("file")
					fmt.Printf("Filepath %q\n", filepath)
					filepath = getPath(filepath)

					http.Handle("/", http.FileServer(http.Dir(filepath)))
					http.ListenAndServe(":4000", nil)

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "Specify index.html file",
						Required: true,
					},
					// TODO: --port flag
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
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
func getPath(path string) string {
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

		// checks if index.html exsists in given path
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			log.Fatal("index.html not found")
		}
	}
	fmt.Println(resultPath)
	return resultPath
}

// checks if program is running on Windows os
func isWindows() bool {
	return runtime.GOOS == "windows"
}
