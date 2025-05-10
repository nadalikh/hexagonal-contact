package app_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"graph/app"
	"graph/domain"
	"graph/dto"
	"graph/internal"
	"graph/logger"
	"graph/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDb *gorm.DB

func TestMain(m *testing.M) {

	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env file")
	}
	var err error
	host := os.Getenv("POSTGRES_SERVER")
	port := os.Getenv("POSTGRES_PORT")
	DBName := os.Getenv("POSTGRES_DB")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	testDb, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", host, username, password, DBName, port)), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}
	_ = testDb.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	_ = testDb.AutoMigrate(&domain.Contact{}, &domain.PhoneNumber{})
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("phone", internal.PhoneValidation); err != nil {
			logger.Error(err.Error())
		}
	}

	code := m.Run()
	os.Exit(code)
}

func setupRouterWithTestDB() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	contactHandler := app.NewContactRestHandler(service.NewContactService(domain.NewContactPostgresRepo(testDb)))
	r.POST("/contact", contactHandler.CreateOne)
	r.PUT("/contact/update", contactHandler.Update)

	return r
}

func TestCreateContact(t *testing.T) {

	router := setupRouterWithTestDB()

	body := map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"phone_numbers": []map[string]string{
			{"number": "09391234567"},
		},
	}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/contact", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	fmt.Println("Create Response:", resp.Body.String())
}

func TestUpdateContact(t *testing.T) {
	// create contact first using repo
	repo := domain.NewContactPostgresRepo(testDb)

	contact := dto.ContactRequestDto{
		FirstName: "Ali",
		LastName:  "Testi",
		PhoneNumbers: []dto.PhoneNumberRequestDto{
			{Number: "09391112222"},
		},
	}
	con, err := repo.Create(contact)

	assert.Nil(t, err)

	router := setupRouterWithTestDB()

	updateBody := map[string]interface{}{
		"contact_id": con.ID,
		"first_name": "Updated",
		"last_name":  "Person",
		"phone_numbers": []map[string]string{
			{
				"phone_id": con.PhoneNumbers[0].ID,
				"number":   "09129998877",
			},
		},
	}

	jsonUpdate, _ := json.Marshal(updateBody)

	req, _ := http.NewRequest(http.MethodPut, "/contact/update", bytes.NewBuffer(jsonUpdate))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println("sjflskjf;ls", resp.Code)

	assert.Equal(t, http.StatusOK, resp.Code)
	fmt.Println("Update Response:", resp.Body.String())
}
