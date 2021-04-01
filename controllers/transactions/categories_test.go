package transactions

import (
	"backend/models/transactions"
	"backend/utils"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

var (
	// this token will expire in 29/3/2022
	oneYearExpiringToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDg1NTk5ODAsImlkIjoyN30.sa61H9Z9ovIwvlVjiV9i45KDZl-XwJyElZpaaY9NJWw"

	categoryJSONRequest = map[string]interface{}{
		"title": "Test Category",
	}
	categoriesURL = "/auth/categories"

	userClaims = &jwt.Token{
		Claims: jwt.MapClaims{
			"exp": 1648559980,
			"id": float64(27),
		},
	}
	e *echo.Echo
	createdID uint
)

type categoryJSONResponse struct {
	Category transactions.Category `json:"category"`
}

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file, %v", err)
	}

	e = echo.New()
	customValidator := &utils.CustomValidator{Validator: validator.New()}
	customValidator.TranslateErrors()
	e.Validator = customValidator

	os.Exit(m.Run())
}


func TestGetCategories(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, categoriesURL, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.Set("user", userClaims)

	if assert.NoError(t, GetCategories(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateCategory(t *testing.T) {
	str, _ := json.Marshal(categoryJSONRequest)
	req := httptest.NewRequest(http.MethodPost, categoriesURL, strings.NewReader(string(str)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", userClaims)

	if assert.NoError(t, CreateCategory(c)) {
		var category categoryJSONResponse
		_ = json.Unmarshal([]byte(rec.Body.String()), &category)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, categoryJSONRequest["title"], category.Category.Title)
		createdID = category.Category.ID
	}
}

func TestDeleteCategory(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, categoriesURL, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("categoryID")
	c.SetParamValues(strconv.Itoa(int(createdID)))

	c.Set("user", userClaims)

	if assert.NoError(t, DeleteCategory(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestDeleteNonExistingCategory(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, categoriesURL, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("categoryID")
	c.SetParamValues("5842041")

	c.Set("user", userClaims)

	if assert.NoError(t, DeleteCategory(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}
