package main

import (
	utils "github.com/gnulinuxindia/internet-chowkidar/clientutils"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

func cliInit(conf, data string) *cli.App {
	updateConf := false
	stopSync := false
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
		Action: func(cCtx *cli.Context) error {
			db, err := utils.FindDatabase(cCtx.String("database"))
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}
			defer db.Close()
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt)
			go func() {
				<-done
				db.Close()
				os.Exit(0)
			}()
			err = utils.Run(cCtx.String("config"), db, &updateConf, &stopSync)
			return err
		},
		Commands: []*cli.Command{
			{
				Name:   "setup",
				Usage:  "setup the chowkidar daemon (interactive)",
				Action: Setup,
			},
		},
	}

	app.Suggest = true

	return app
}
