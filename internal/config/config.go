package config

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/gnulinuxindia/internet-chowkidar/utils"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	// used by the repositories (required)
	DatabaseURL    string `koanf:"DATABASE_URL"`
	DatabaseDriver string `koanf:"DATABASE_DRIVER"`

	// used by tracing
	ServiceName    string `koanf:"SERVICE_NAME"`
	ServiceVersion string `koanf:"SERVICE_VERSION"`

	// general
	Env string `koanf:"ENV"`
}

var Conf *Config

var k = koanf.New(".")

var allowedDBTypes = []string{
	"sqlite3",
	"mysql",
}

var defaultConfigProvider = confmap.Provider(map[string]any{
	"DATABASE_URL":    "file:sqlite.db?mode=rwc&cache=shared&_fk=1",
	"DATABASE_DRIVER": "sqlite3",
	"SERVICE_NAME":    "go-template",
	"SERVICE_VERSION": "0.1.0",
	"ENV":             "local",
}, "")

func ProvideConfig() (*Config, error) {
	if Conf == nil {
		err := k.Load(defaultConfigProvider, nil)
		if err != nil {
			slog.Error("error loading default config", "err", err)
			os.Exit(1)
		}

		_ = k.Load(file.Provider(".env"), dotenv.Parser())

		if err := k.Load(env.Provider("", ".", nil), nil); err != nil {
			slog.Error("error loading config", "err", err)
			os.Exit(1)
		}

		Conf = &Config{}
		err = k.Unmarshal("", Conf)
		if err != nil {
			slog.Error("error loading config", "err", err)
			os.Exit(1)
		}
	}

	if !slices.Contains(allowedDBTypes, Conf.DatabaseDriver) {
		return nil, unknownOptionError("database driver", Conf.DatabaseDriver, allowedDBTypes)
	}

	return Conf, nil
}

func unknownOptionError(valueType string, providedValue string, allowedValues []string) error {
	return fmt.Errorf("unknown %s provided to config: '%s'\nmust be one of: %s", valueType, providedValue,
		strings.Join(
			utils.MapOver(allowedValues, func(s string) string { return fmt.Sprintf("'%s'", s) }),
			", ",
		))
}
