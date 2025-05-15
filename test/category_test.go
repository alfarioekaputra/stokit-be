package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"stokit/internal/entity"
	"stokit/internal/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCategoryWithoutParent(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("email = ?", "admin@mail.com").First(user).Error
	assert.Nil(t, err)

	requestBody := model.CreateCategoryRequest{
		Name:     "Makanan",
		ParentID: "",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/category", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
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
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.ParentID, "")
	assert.NotNil(t, responseBody.Data.ID)
}

func TestCreateCategoryWithParent(t *testing.T) {
	TestLogin(t)
	TestCreateCategoryWithoutParent(t)

	user := new(entity.User)
	err := db.Where("email = ?", "admin@mail.com").First(user).Error
	assert.Nil(t, err)

	category := new(entity.Category)
	err2 := db.Where("name = ?", "Makanan").First(category).Error
	assert.Nil(t, err2)

	requestBody := model.CreateCategoryRequest{
		Name:     "Mie Instant",
		ParentID: category.ID,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/category", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
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
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.ParentID, responseBody.Data.ParentID)
	assert.NotNil(t, responseBody.Data.ID)
}

func TestViewCategoryWithoutParentById(t *testing.T) {
	TestLogin(t)
	TestCreateCategoryWithoutParent(t)

	user := new(entity.User)
	err := db.Where("email = ?", "admin@mail.com").First(user).Error
	assert.Nil(t, err)

	category := new(entity.Category)
	err2 := db.Where("name = ?", "Makanan").First(category).Error
	assert.Nil(t, err2)

	request := httptest.NewRequest(http.MethodGet, "/api/category/"+category.ID+"/view", nil)
	request.Header.Set("Content-Type", "application/json")
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
	assert.Equal(t, category.ID, responseBody.Data.ID)
	assert.Equal(t, category.Name, responseBody.Data.Name)
	assert.Equal(t, category.ParentID, "")
}
