package main

import (
	"golang-blog-api/api"
	"golang-blog-api/config"
	"golang-blog-api/db"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	router := api.Routes()

	if err := config.InitConfig(); err != nil {
		log.Println("Failed to read config.yaml: ", err)
	}

	if err := http.ListenAndServe(viper.GetString("port"), router); err != nil {
		log.Println("Не удалось запустить сервер: ", err)
		return
	}
	log.Println("Сервер запущен")
}
