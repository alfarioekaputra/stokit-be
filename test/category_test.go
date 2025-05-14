package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"stokit/internal/entity"
	"stokit/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCategory(t *testing.T) {
	user := new(entity.User)
	err := db.Where("username = ?", "admin").First(user).Error
	assert.Nil(t, err)

	category := new(entity.Category)
	err = db.Find(&category).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodGet, "/api/category", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CategoryResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, category.Name, responseBody.Data.Name)
	assert.NotNil(t, responseBody.Data.ID)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
}
