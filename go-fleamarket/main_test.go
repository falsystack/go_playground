package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"go-fleamarket/dto"
	"go-fleamarket/infra"
	"go-fleamarket/models"
	"go-fleamarket/services"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestMain 은 다른 테스트 함수가 실행되기 전에 먼저 실행된다, beforeAll 느낌
func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{
			Name:        "테스트 아이템1",
			Price:       1000,
			Description: "",
			SoldOut:     false,
			UserID:      1,
		},
		{
			Name:        "테스트 아이템2",
			Price:       2000,
			Description: "테스트2",
			SoldOut:     true,
			UserID:      1,
		},
		{
			Name:        "테스트 아이템3",
			Price:       3000,
			Description: "테스트3",
			SoldOut:     false,
			UserID:      2,
		},
	}

	users := []models.User{
		{Email: "test1@test.com", Password: "123456"},
		{Email: "test2@test.com", Password: "234567"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	return setupRouter(db)
}

func TestFindAll(t *testing.T) {
	// test setup
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/items", nil)

	// Run API request
	router.ServeHTTP(w, req)

	// Get API Response
	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	// test setup
	router := setup()
	token, err := services.CreateToken(1, "test1@test.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "테스트 아이템4",
		Price:       4000,
		Description: "Create test",
	}
	// object -> json
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	// Run API request
	router.ServeHTTP(w, req)

	// Get API Response
	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}

func TestCreateUnauthorized(t *testing.T) {
	// test setup
	router := setup()

	createItemInput := dto.CreateItemInput{
		Name:        "테스트 아이템4",
		Price:       4000,
		Description: "Create test",
	}
	// object -> json
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(reqBody))

	// Run API request
	router.ServeHTTP(w, req)

	// Get API Response
	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
