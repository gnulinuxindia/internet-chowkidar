package main

import (
	"log"
	"os"

	"github.com/gnulinuxindia/internet-chowkidar/cmd/chowkidar/utils"
)

func main() {
	conf, data := utils.DefaultStoragePath()
	app := cliInit(conf, data)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
