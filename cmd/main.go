package main

import (
	"golang-blog-api/db"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Println("Не удалось запустить сервер: ", err)
		return
	} else {
		log.Println("Сервер запущен")
	}
}
