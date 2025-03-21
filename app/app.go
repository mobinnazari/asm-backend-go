package app

import (
	"log"
	"net/http"
	"strings"

	"git.sindadsec.ir/asm/backend/config"
	"git.sindadsec.ir/asm/backend/db"
	"git.sindadsec.ir/asm/backend/docs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	DB     *gorm.DB
	Logger *zap.SugaredLogger
}

func Init(config *config.Config) *Application {
	db := db.Init(config.MysqlAddr)
	logger := zap.Must(zap.NewProduction()).Sugar()

	app := &Application{
		Config: config,
		DB:     db,
		Logger: logger,
	}
	return app
}

func (app *Application) Run() {
	docs.SwaggerInfo.Schemes = []string{strings.Split(app.Config.ApiUrl, "://")[0]}
	docs.SwaggerInfo.Host = strings.Split(app.Config.ApiUrl, "://")[1]

	srv := &http.Server{
		Addr:    app.Config.Addr,
		Handler: app.mount(),
	}

	app.Logger.Infow("server has started", "addr", app.Config.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Panicln(err.Error())
	}
}
