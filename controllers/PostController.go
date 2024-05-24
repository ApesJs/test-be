package controllers

import (
	"github.com/ApesJs/test-be/helpers"
	"github.com/ApesJs/test-be/initializers"
	"github.com/ApesJs/test-be/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var PostRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `gorm:"type:TEXT" json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

var validate = validator.New()

func CreatePost(context *gin.Context) {
	err := context.Bind(&PostRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request data from body"})
		return
	}

	post := models.Post{
		Title:    PostRequest.Title,
		Content:  PostRequest.Content,
		Category: PostRequest.Category,
		Status:   PostRequest.Status,
	}

	validateErr := validate.Struct(post)
	if validateErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create Post: " + validateErr.Error()})
		return
	}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Create User: " + result.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Post Created"})

}

func FindAllPosts(context *gin.Context) {
	perPage := 6
	page := 1
	pageStr := context.Param("page")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	pagination := helpers.GetPaginationData(page, perPage, models.Post{})

	var posts []models.Post
	initializers.DB.Limit(perPage).Offset(pagination.Offset).Find(&posts)

	context.JSON(http.StatusOK, gin.H{
		"data_posts": posts,
		"pagination": pagination,
	})
}

func FindByIDPost(context *gin.Context) {
	postID := context.Param("id")

	var post models.Post
	initializers.DB.First(&post, postID)

	context.JSON(http.StatusOK, post)
}

func UpdatePost(context *gin.Context) {
	postID := context.Param("id")

	err := context.Bind(&PostRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request data from body"})
		return
	}

	var post models.Post
	initializers.DB.First(&post, postID)

	initializers.DB.Model(&post).Updates(models.Post{
		Title:    PostRequest.Title,
		Content:  PostRequest.Content,
		Category: PostRequest.Category,
		Status:   PostRequest.Status,
	})

	validateErr := validate.Struct(PostRequest)
	if validateErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Update Post: " + validateErr.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post Updated"})
}

func DeletePost(context *gin.Context) {
	postID := context.Param("id")

	initializers.DB.Delete(&models.Post{}, postID)

	context.JSON(http.StatusOK, gin.H{"message": "Delete Success"})
}

func TrashedPost(context *gin.Context) {
	postID := context.Param("id")

	var post models.Post
	initializers.DB.First(&post, postID)

	initializers.DB.Model(&post).Updates(models.Post{
		Status: "trash",
	})

	context.JSON(http.StatusOK, gin.H{"message": "Post Trashed Success"})
}
