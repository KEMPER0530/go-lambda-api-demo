package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"log"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// MySQL用ドライバ
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm"
		"go-lambda-api-demo/src/infrastructure"
)

func TestMain(m *testing.M) {
	fmt.Println("before test main_test.go")

	// 環境変数ファイルの読込
	if os.Getenv("GO_ENV") == "development" {
		err := godotenv.Load(fmt.Sprintf("../src/infrastructure/%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal(err)
		}
	}
	code := m.Run()
	fmt.Println("after test serve_test.go")
	os.Exit(code)
}

func TestActuaterHealth(t *testing.T) {
	t.Log("START TestActuatorHealth")

	gin.SetMode(gin.TestMode)
	r := infrastructure.NewRouting()
	r.Gin.Use(gin.Logger())

	req, _ := http.NewRequest("GET", "/V1/actuator-health", nil)
	rec := httptest.NewRecorder()
	r.Gin.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	t.Log("END TestActuatorHealth")
}
