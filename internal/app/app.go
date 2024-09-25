package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fleeper2133/tasks-app/internal/handler"
	"github.com/fleeper2133/tasks-app/internal/pkg"
	"github.com/fleeper2133/tasks-app/internal/repository"
	"github.com/fleeper2133/tasks-app/internal/server"
	"github.com/fleeper2133/tasks-app/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

const (
	configName = "config"
)

func Run(configPath string) {
	if err := initConfig(configPath); err != nil {
		logrus.Fatalf("Error init configs: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error init .env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Fail to init db: %s", err)
	}

	jwtManager := pkg.NewTokenJWTManager()

	mailManager := pkg.NewSendMailManager(pkg.EmailConfig{
		From:     viper.GetString("mail.from"),
		Password: viper.GetString("mail.password"),
		SmtpHost: viper.GetString("mail.host"),
		SmtpPort: viper.GetString("mail.port"),
	})

	repo := repository.NewRepository(db)
	service := service.NewService(repo, jwtManager, mailManager)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRouter()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("TaskApp Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TaskApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection: %s", err.Error())
	}
}

func initConfig(configPath string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
