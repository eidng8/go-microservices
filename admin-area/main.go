package main

import (
	"log"
	"os"

	"eidng8.cc/microservices/admin-area/ent"
	_ "eidng8.cc/microservices/admin-area/ent/runtime"
	"eidng8.cc/microservices/rdbms/mysql"
)

type Env struct {
	db *ent.Client
}

func main() {
	env := setup()
	migrate()
	start(env)
}

func setup() *Env {
	host, user, pass, dbname := getDbCfg()
	cfg := mysql.GetConnCfg(host, user, pass, dbname)
	db, err := ent.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}
	return &Env{db: db}
}

func start(env *Env) {
	r := setupRouter(env)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func migrate() {
	host, user, pass, dbname := getDbCfg()
	mysql.Migrate(host, user, pass, dbname)
}

func getDbCfg() (string, string, string, string) {
	host := os.Getenv("DB_CONN") // unix(/var/run/mysqld/mysql.sock)
	if "" == host {
		log.Fatal("DB_CONN is not set")
	}
	user := os.Getenv("DB_USER") // root
	if "" == host {
		log.Fatal("DB_USER is not set")
	}
	pass := os.Getenv("DB_PASSWORD") // 123456
	if "" == pass {
		log.Fatal("DB_PASSWORD is not set")
	}
	dbname := os.Getenv("DB_NAME") // admin_areas
	if "" == dbname {
		log.Fatal("DB_NAME is not set")
	}
	return host, user, pass, dbname
}
