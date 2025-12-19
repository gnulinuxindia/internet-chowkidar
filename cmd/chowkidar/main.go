package main

import (
	"log"
	"os"

	utils "github.com/gnulinuxindia/internet-chowkidar/clientutils"
)

func main() {
	conf, data := utils.DefaultStoragePath()
	app := cliInit(conf, data)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
