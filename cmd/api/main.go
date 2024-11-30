package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/linkeunid/api.linkeun.com/internal/version"
	"github.com/linkeunid/api.linkeun.com/pkg/bcrypt"
	"github.com/linkeunid/api.linkeun.com/pkg/config"
	"github.com/linkeunid/api.linkeun.com/pkg/logger"
	"github.com/linkeunid/api.linkeun.com/pkg/sentry"
)

func main() {
	logger := logger.InitLogger()

	err := bootstrap(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type App struct {
	config *config.Config
	// db     *database.DB
	logger *slog.Logger
	bcrypt *bcrypt.Bcrypt
	wg     sync.WaitGroup
}

func bootstrap(logger *slog.Logger) error {
	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("Version: %s\n", version.Get())
		return nil
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// db, err := database.NewDB(cfg.Dsn)
	// if err != nil {
	// 	return err
	// }
	// defer func() {
	// 	sqlDB, err := db.DB.DB()
	// 	if err != nil {
	// 		logger.Error("failed to get sql.DB", "error", err)
	// 	} else {
	// 		sqlDB.Close()
	// 	}
	// }()

	err = sentry.InitSentry(cfg.SentryDsn)
	if err != nil {
		return err
	}

	bcrypt := bcrypt.NewBcrypt(cfg.AppSalt)

	app := &App{
		logger: logger,
		// db:     db,
		config: cfg,
		bcrypt: bcrypt,
	}

	return app.serveHTTP()
}
