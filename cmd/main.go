package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/gavrylenkoIvan/todo-app-gin"
	"github.com/gavrylenkoIvan/todo-app-gin/pkg/handler"
	"github.com/gavrylenkoIvan/todo-app-gin/pkg/repository"
	"github.com/gavrylenkoIvan/todo-app-gin/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error occured while initializing server, error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error occured while initializing server, error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Error occured while initializing server, %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while initializing server, %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
