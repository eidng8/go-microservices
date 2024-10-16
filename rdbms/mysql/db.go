package mysql

import (
	"database/sql"
	"time"

	ms "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connect(addr, user, password, dbname string) (*sql.DB, error) {
	cfg := GetConnCfg(addr, user, password, dbname)
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

func GetConnCfg(addr, user, password, dbname string) ms.Config {
	return ms.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbname,
		MultiStatements:      true,
		AllowNativePasswords: true,
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  time.Local,
	}
}
