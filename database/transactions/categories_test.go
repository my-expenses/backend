package transactions

import (
	dbInstance "backend/database"
	customErrors "backend/database/errors"
	categoriesModel "backend/models/transactions"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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
	var category categoriesModel.Category
	err := dbInstance.GetDBConnection().Where("user_id = 27 AND title = 'Testing Service Layer'").First(&category).Error
	assert.Equal(t, category.ID, uint(0))
	assert.Error(t, err, gorm.ErrRecordNotFound)

	category = categoriesModel.Category{
		UserID:      27,
		Title:       "Testing Service Layer",
	}
	err = CreateCategory(&category)
	createdCategoryID = category.ID
	assert.Equal(t, err, nil)
	assert.Equal(t, category.UserID, uint(27))
	assert.Equal(t, category.Title, "Testing Service Layer")

	category = categoriesModel.Category{}
	err = dbInstance.GetDBConnection().Where("user_id = 27 AND title = 'Testing Service Layer'").First(&category).Error
	assert.NotEqual(t, category.ID, uint(0))
	assert.Equal(t, err, nil)
}

func TestGetMoreThanZeroCategories(t *testing.T) {
	categories := GetCategories(27)
	assert.Greater(t, len(categories), 0)
}

func TestCreateCategoryConflict(t *testing.T) {
	category := categoriesModel.Category{
		UserID:      27,
		Title:       "Testing Service Layer",
	}
	err := CreateCategory(&category)
	assert.Equal(t, err, &customErrors.DuplicateCategoryError{})
}

func TestUpdateCategory(t *testing.T) {
	category := categoriesModel.Category{
		UserID: 27,
		Title: "Updated category",
	}
	category.ID = createdCategoryID
	err := UpdateCategory(&category)
	assert.Equal(t, nil, err)

	err = dbInstance.GetDBConnection().Where("user_id = 27 AND title = 'Updated category'").First(&category).Error
	assert.Equal(t, category.UserID, uint(27))
	assert.Equal(t, category.Title, "Updated category")
}

func TestDeleteCategory(t *testing.T) {
	err := DeleteCategory(createdCategoryID, 27)
	assert.Equal(t, err, nil)
}

func TestDeleteNoCategoryFound(t *testing.T) {
	err := DeleteCategory(521569, 27)
	assert.Equal(t, err, &customErrors.RecordNotFoundError{})
}
