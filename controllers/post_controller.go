package controllers

import (
    "mini-social-media/services"
    "mini-social-media/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)


func CreatePost(c *gin.Context) {
	var request struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	post, err := services.CreatePost(request.Content)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Could not create post")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	post, err := services.UpdatePost(id, request.Content)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, post)
}

func GetAllPosts(c *gin.Context) {
	posts := services.GetAllPosts()
	utils.SuccessResponse(c, http.StatusOK, posts)
}

func LikePost(c *gin.Context) {
	id := c.Param("id")

	post, err := services.LikePost(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, post)
}

func AddComment(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Comment string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	post, err := services.AddComment(id, request.Comment)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, post)
}

func GetPostDetails(c *gin.Context) {
	id := c.Param("id")

	post, err := services.GetPostDetails(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, post)
}
