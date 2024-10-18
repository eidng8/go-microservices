package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var env *Env
var router *gin.Engine

func setupTest(tb testing.TB) func(tb testing.TB) {
	env = setup()
	tx, err := env.db.BeginTx(context.Background(), nil)
	if err != nil {
		tb.Fatalf("Could not start transaction: %v", err)
	}
	router = setupRouter(env)
	return func(tb testing.TB) {
		if err := tx.Rollback(); err != nil {
			tb.Fatalf("Could not rollback transaction: %v", err)
		}
		if err = env.db.Close(); err != nil {
			tb.Fatalf("Could not close db: %v", err)
		}
	}
}

func Test_create(t *testing.T) {
	router := setupRouter(nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func Test_detail(t *testing.T) {
	router := setupRouter(nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func Test_list(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	q := req.URL.Query()
	q.Set("page", "1")
	q.Set("per_page", "5")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func Test_remove(t *testing.T) {
	router := setupRouter(nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func Test_restore(t *testing.T) {
	router := setupRouter(nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func Test_update(t *testing.T) {
	router := setupRouter(nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
