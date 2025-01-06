package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	handler "github.com/Njrctr/DeNet_test/internal/handlers"
	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/Njrctr/DeNet_test/internal/repository"
	"github.com/Njrctr/DeNet_test/internal/repository/postgres"
	"github.com/Njrctr/DeNet_test/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Wallet API
// @version 1.0
// @description API Server for Wallet

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Run(configPath string) {
	if err := initConfig(configPath); err != nil {
		logrus.Fatalf("Ошибка инициализации конфига: %v", err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Ошибка получения переменных окружения: %s", err.Error())
	}
	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Ошибка инициализации Базы данных: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(models.Server)
	logrus.Printf("Попытка запуска сервера на порту %s", viper.GetString("port"))

	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRouters()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("DeNet Backend Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("DeNet Backend Stoped")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig(confPath string) error {
	viper.AddConfigPath(confPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
