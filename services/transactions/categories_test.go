package transactions

import (
	customErrors "backend/database/errors"
	categoriesModel "backend/models/transactions"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var createdCategoryID uint

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file, %v", err)
	}
	os.Exit(m.Run())
}

func TestCreateCategory(t *testing.T) {
	category := categoriesModel.Category{
		UserID:      27,
		Title:       "Testing Service Layer",
		Description: "Creating category from service layer test",
	}
	err := CreateCategory(&category)
	createdCategoryID = category.ID
	assert.Equal(t, err, nil)
}

func TestCreateCategoryConflict(t *testing.T) {
	category := categoriesModel.Category{
		UserID:      27,
		Title:       "Testing Service Layer",
		Description: "Creating category from service layer test",
	}
	err := CreateCategory(&category)
	assert.Equal(t, err, &customErrors.DuplicateCategoryError{})
}

func TestDeleteCategory(t *testing.T) {
	err := DeleteCategory(createdCategoryID, 27)
	assert.Equal(t, err, nil)
}

func TestDeleteNoCategoryFound(t *testing.T) {
	err := DeleteCategory(521569, 27)
	assert.Equal(t, err, &customErrors.RecordNotFoundError{})
}