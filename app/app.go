package app

import (
	"log"
	"net/http"

	"git.sindadsec.ir/asm/backend/config"
	"go.uber.org/zap"
)

type Application struct {
	Config *config.Config
	Logger *zap.SugaredLogger
}

func Init(config *config.Config) *Application {
	logger := zap.Must(zap.NewProduction()).Sugar()
	app := &Application{
		Config: config,
		Logger: logger,
	}
	return app
}

func (app *Application) Run() {
	srv := &http.Server{
		Addr:    app.Config.Addr,
		Handler: app.mount(),
	}

	app.Logger.Infow("server has started", "addr", app.Config.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Panicln(err.Error())
	}
}
