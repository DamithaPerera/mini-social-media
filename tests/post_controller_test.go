package tests

import (
	"bytes"
	"encoding/json"
	"mini-social-media/models"
	"mini-social-media/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.RegisterPostRoutes(router)
	return router
}

func TestCreatePost(t *testing.T) {
	router := setupRouter()
	requestBody := map[string]string{"content": "Hello, World!"}
	body, _ := json.Marshal(requestBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/posts/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "success", response["status"])
	assert.NotEmpty(t, response["data"].(map[string]interface{})["id"])
}

func TestGetAllPosts(t *testing.T) {
	router := setupRouter()

	// Preload data
	models.PostsLock.Lock()
	models.Posts["1"] = &models.Post{ID: "1", Content: "Post 1"}
	models.PostsLock.Unlock()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts/", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "success", response["status"])
	assert.NotEmpty(t, response["data"])
}

func TestLikePost(t *testing.T) {
	router := setupRouter()

	// Preload data
	models.PostsLock.Lock()
	models.Posts["1"] = &models.Post{ID: "1", Content: "Post 1"}
	models.PostsLock.Unlock()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/posts/1/like", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "success", response["status"])
	assert.Equal(t, float64(1), response["data"].(map[string]interface{})["likes"])
}