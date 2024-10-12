package main

import (
	"log"
	"os"

	"eidng8.cc/microservices/rdbms/mysql"
)

var host = os.Getenv("DB_CONN")     // unix(/var/run/mysqld/mysql.sock)
var user = os.Getenv("DB_USER")     // root
var pass = os.Getenv("DB_PASSWORD") // 123456
var dbname = os.Getenv("DB_NAME")   // admin_areas

func main() {
	migrate()
	start()
}

func migrate() {
	mysql.Migrate(host, user, pass, dbname)
}

func start() {
	r := setupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
