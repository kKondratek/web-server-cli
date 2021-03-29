package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

func getPath() {
	// TODO: validate --file arg and return proper filepath
}
