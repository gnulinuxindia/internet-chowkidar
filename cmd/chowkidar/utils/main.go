package utils

import (
	"log"
	"runtime/debug"

	"github.com/urfave/cli/v2"
)

func Version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value[:8]
			}
		}
	}
	return "unknown, please build with Go 1.13+ or use Git"
}

func Debug(text string, cCtx *cli.Context) {
	if cCtx.Bool("debug") {
		log.Println("[DEBUG]: " + text)
	}
}
