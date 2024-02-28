package db

import (
	"database/sql"
	"fmt"
	"golang-blog-api/config"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

var DB *sql.DB

type dbConfig struct {
	user     string
	password string
	port     string
	host     string
	dbname   string
	sslmode  string
}

func initDSN() string {
	if err := config.InitConfig(); err != nil {
		log.Println("Failed to read config.yaml: ", err)
	}

	dbConfig := dbConfig{
		user:     viper.GetString("database.user"),
		password: viper.GetString("database.password"),
		port:     viper.GetString("database.port"),
		host:     viper.GetString("database.host"),
		dbname:   viper.GetString("database.dbname"),
		sslmode:  viper.GetString("database.sslmode"),
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		dbConfig.user, dbConfig.password, dbConfig.host, dbConfig.port, dbConfig.dbname, dbConfig.sslmode)

	return dsn
}

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", initDSN())
	if err != nil {
		log.Println("Failed to connect: ", err)
		return
	} else {
		log.Println("Connected to DB")
	}

	if err = DB.Ping(); err != nil {
		log.Println("Error connecting to the database:", err)
		return
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Connection was closed")
	}
}
