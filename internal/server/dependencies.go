package server

import (
	stdLog "log"

	"template/internal/clients/db"

	"github.com/go-chi/chi"

	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/config"
	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/log"
	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/platform/router"
)

func SetupDependencies() (chi.Router, config.Configuration) {
	//setea el log
	logger := setupLogger()
	//obtiene la configuracion
	cfg := config.ResolveConfiguration(logger)
	logger.Info("initializing server dependencies")
	//inicializo las rutas
	r := router.NewRouter(router.Base{Logger: logger, Configuration: cfg})
	logger.Info("routes initialized")
	// inicializo la base de datos
	err := db.InitializeDB(cfg)
	if err != nil {
		logger.Info(err.Error())
	}

	logger.Info("Db initialized")

	setupHandlers(r, cfg)
	router.PrintRoutes(logger, r)

	return r, cfg
}

func setupLogger() log.Logger {
	logger, err := log.NewLogger()
	if err != nil {
		stdLog.Printf("defaulting to no-op logger, an error has occurred: %v", err.Error())
		logger = log.DefaultLogger()
	}

	return logger
}
