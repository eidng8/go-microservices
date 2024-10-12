package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//applyMigrate()
	startServer()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET(
		"/ping", func(c *gin.Context) {
			c.JSON(
				http.StatusOK, gin.H{
					"message": "pong",
				},
			)
		},
	)
	return r
}

func startServer() {
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

/*
func applyMigrate() {
	dbu := os.Getenv("DB_USER")
	dbp := os.Getenv("DB_PASSWORD")
	dbn := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@/%s?multiStatements=true", dbu, dbp, dbn)
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Could not apply migrations: %v", err)
	}

}
*/
