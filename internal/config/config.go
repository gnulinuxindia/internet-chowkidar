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
	ServiceName     string `koanf:"SERVICE_NAME"`
	ServiceVersion  string `koanf:"SERVICE_VERSION"`
	TracingExporter string `koanf:"TRACING_EXPORTER"`
	TracingHost     string `koanf:"TRACING_HOST"`
	TracingPort     string `koanf:"TRACING_PORT"`

	// general
	Env    string `koanf:"ENV"`
	Port   string `koanf:"PORT"`
	Listen string `koanf:"LISTEN"`

	ApiKey string `koanf:"API_KEY"`
}

var Conf *Config

var k = koanf.New(".")

var allowedDBTypes = []string{
	"sqlite3",
	"postgres",
}

var defaultConfigProvider = confmap.Provider(map[string]any{
	"DATABASE_URL":     "file:sqlite.db",
	"DATABASE_DRIVER":  "sqlite3",
	"SERVICE_NAME":     "internet-chowkidar",
	"SERVICE_VERSION":  "0.1.0",
	"TRACING_EXPORTER": "http",
	"ENV":              "debug",
	"PORT":             "9000",
	"LISTEN":           "0.0.0.0",
	"TRACING_HOST":     "localhost",
	"TRACING_PORT":     "4318",
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
