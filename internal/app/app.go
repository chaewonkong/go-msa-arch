package app

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

// DeferFunc defer function
type DeferFunc func()

// RunApplication run application and handle signals
//
//	Use this function in main functions to run application
func RunApplication(serviceName string, appFactory Factory) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT) // SIGTERM: when k8s stop pod

	// global logger
	cleanUpLogger := initGlobalLogger(serviceName)
	defer cleanUpLogger()

	//config
	config := appFactory.NewConfig()
	err := ReadConfig(config)
	if err != nil {
		zap.L().Fatal("failed to read config", zap.Error(err))
	}

	// Application init
	app, err := appFactory.NewApp(config)
	if err != nil {
		zap.L().Fatal("failed to create application", zap.Error(err))
	}

	// Run
	err = app.Run()
	if err != nil {
		zap.L().Fatal("application run failed", zap.Error(err))
	}

	// Stop
	<-c
	err = app.Stop()
	if err != nil {
		zap.L().Fatal("application stop failed", zap.Error(err))
	}

	zap.L().Info("application stopped successfully")
}

func initGlobalLogger(serviceName string) DeferFunc {
	logger := zap.
		Must(zap.NewProduction()).
		With(zap.String("service", serviceName)) // add service name to logger
	undo := zap.ReplaceGlobals(logger)

	return func() {
		_ = logger.Sync()
		undo()
	}
}

// App app
type App interface {
	Run() error
	Stop() error
}

// Factory app factory
type Factory interface {
	NewApp(config any) (App, error)
	NewConfig() any
}
