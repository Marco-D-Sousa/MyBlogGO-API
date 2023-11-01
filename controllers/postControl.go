package controllers

import (
	"github.com/gin-gonic/gin"
	"myblog/initializers"
	"myblog/models"
)

func PostCreate(ctx *gin.Context) {
	var body struct {
		Title   string
		Content string
	}

	ctx.Bind(&body)

	post := models.Post{Title: body.Title, Content: body.Content}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(200, gin.H{
			"New post": post,
		})
	}
}

func PostIndex(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	ctx.JSON(200, gin.H{
		"All Posts": posts,
	})
}

func PostShow(ctx *gin.Context) {
	id := ctx.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	ctx.JSON(200, gin.H{
		"Post": post,
	})
}

func PostUpdate(ctx *gin.Context) {

}
