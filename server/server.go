package server

import (
	"go.uber.org/zap"

	"github.com/tiennv147/http-echo/config"
)

// Run ...
func Run(f *config.Flags) {
	// Use a temporary logger to parse the configuration and output.
	tmpLogger := logging.NewTmpLogger().With(zap.String("filename", f.ConfigPath))

	var cfg chedwig.Config
	if err := config.ParseFile(f.ConfigPath, &cfg, f.Template); err != nil {
		tmpLogger.Fatal("parsing configuration failed", zap.Error(err))
	}

	RunWithConfig(&cfg)
}

// RunWithConfig ...
func RunWithConfig(cfg *chedwig.Config) {
	// Init the logger.
	logger, err := logging.NewLogger(cfg.Logger)
	if err != nil {
		logging.NewTmpLogger().Fatal("could not instantiate logger", zap.Error(err))
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			panic(err)
		}
	}()

	// Lock by Serve
}
