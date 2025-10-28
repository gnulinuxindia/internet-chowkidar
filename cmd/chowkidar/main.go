package main

import (
	"log"
	"os"

	"github.com/gnulinuxindia/internet-chowkidar/cmd/chowkidar/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	conf, data := utils.DefaultStoragePath()
	app := &cli.App{
		Name:    "Internet Chowkidar",
		Usage:   "Run the chowkidar daemon to report blocked sites",
		Version: utils.Version(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   conf,
				Usage:   "config file to read from",
			},
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"d"},
				Value:   data,
				Usage:   "config file to read from",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Value: false,
				Usage: "Change verbosity of the application",
			},
		},
		Action: Run,
		Commands: []*cli.Command{
			{
				Name:   "setup",
				Usage:  "setup the chowkidar daemon (interactive)",
				Action: Setup,
			},
		},
	}

	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
