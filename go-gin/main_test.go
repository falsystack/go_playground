package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go-gin/dto"
	"go-gin/infra"
	"go-gin/models"
	"go-gin/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// 他のテストの前に呼び出される JUnitのBeforeAll
func TestMain(t *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := t.Run() // このファイルに含まれる全てのテスト関数が呼び出される

	os.Exit(code) // テストの終了
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{
			UserID:      1,
			Name:        "テスト商品１",
			Price:       1000,
			Description: "説明１",
			SoldOut:     false,
		},
		{
			UserID:      1,
			Name:        "テスト商品２",
			Price:       2000,
			Description: "説明２",
			SoldOut:     true,
		},
		{
			UserID:      2,
			Name:        "テスト商品３",
			Price:       3000,
			Description: "説明３",
			SoldOut:     false,
		},
	}

	users := []models.User{
		{
			Email:    "testuser1@test.com",
			Password: "1q2w3e4r",
		},
		{
			Email:    "testuser2@test.com",
			Password: "1q2w3e4r",
		},
	}

	for _, item := range items {
		db.Create(&item)
	}
	for _, user := range users {
		db.Create(&user)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	// テストのセットアップ
	router := setup()

	w := httptest.NewRecorder()                      // responseを記録するためのrecorder
	req, _ := http.NewRequest("GET", "/items/", nil) // mock request

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string][]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	// テストのセットアップ
	router := setup()

	token, err := services.CreateToken(1, "testuser1@test.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "テスト商品３",
		Price:       4000,
		Description: "説明4",
	}
	reqBody, _ := json.Marshal(createItemInput)

	// responseを記録するためのrecorder
	w := httptest.NewRecorder()
	// mock request
	req, _ := http.NewRequest("POST", "/items/", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果を取得
	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	// アサーション
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}
