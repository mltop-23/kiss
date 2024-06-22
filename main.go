package main

import (
	"context"
	"kissandeat/cmd/api"
	"kissandeat/cmd/handlers"
	"kissandeat/internal/repository"
	"kissandeat/internal/repository/dbrepo"
	"kissandeat/internal/service"

	// "kissandeat/internal/service/dbrepo"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
func main() {
	secretKey := "your_secret_key"
	refreshSecret := "your_refresh_secret_key"
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error loading env variables: %s", err.Error())
	// }
	db, _, err := dbrepo.NewDB(dbrepo.Config{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DBName:   viper.GetString("db.dbname"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos, err := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(secretKey, refreshSecret, services)
	srv := new(api.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("kissandpizz Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("kissandpizz Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
