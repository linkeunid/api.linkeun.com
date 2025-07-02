package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/goravel/framework/contracts/queue"
	"github.com/goravel/framework/facades"

	"github.com/linkeunid/api.linkeun.com/bootstrap"
)

func testRedisConnection() error {
	appName := facades.Config().GetString("app.name")
	hasAccess := facades.Cache().Has(appName)
	if !hasAccess {
		return errors.New("redis connection failed")
	}

	return nil
}

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route Run error: %v", err)
		}
	}()

	if err := testRedisConnection(); err == nil {
		facades.Log().Infof("Redis connection success")

		// Start queue server by facades.Queue().
		go func() {
			if err := facades.Queue().Worker(queue.Args{
				Queue: "mails",
			}).Run(); err != nil {
				facades.Log().Errorf("Queue run error: %v", err)
			}
		}()
	}

	// Listen for the OS signal
	go func() {
		<-quit
		if err := facades.Route().Shutdown(); err != nil {
			facades.Log().Errorf("Route Shutdown error: %v", err)
		}

		os.Exit(0)
	}()

	select {}
}
