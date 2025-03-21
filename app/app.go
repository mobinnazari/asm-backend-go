package app

import (
	"log"
	"net/http"
	"strings"

	"git.sindadsec.ir/asm/backend/config"
	"git.sindadsec.ir/asm/backend/db"
	"git.sindadsec.ir/asm/backend/docs"
	"git.sindadsec.ir/asm/backend/mail"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Email  *gomail.Dialer
	Logger *zap.SugaredLogger
}

func Init(config *config.Config) *Application {
	mysql := db.Init(config.MysqlAddr)
	redis := db.InitRedis(config.RedisAddr)
	email := mail.Init(config.Mail.Host, config.Mail.Username, config.Mail.Password, config.Mail.Port)
	logger := zap.Must(zap.NewProduction()).Sugar()

	app := &Application{
		Config: config,
		DB:     mysql,
		Redis:  redis,
		Email:  email,
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
