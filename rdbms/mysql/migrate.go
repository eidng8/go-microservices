package mysql

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func makeDsn(host, user, password, dbname string) string {
	return fmt.Sprintf(
		"mysql://%s:%s@tcp(%s)/%s?multiStatements=true",
		user, password, host, dbname,
	)
}

func Migrate(host, user, password, dbname string) {
	dsn := makeDsn(host, user, password, dbname)
	applyMigrate(dsn)
}

func applyMigrate(dsn string) {
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Could not apply migrations: %v", err)
	}
}
