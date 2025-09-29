package controller

import "github.com/gin-gonic/gin"

// CommentController defines the interface for handling comment-related HTTP requests
type CommentController interface {
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	GetCommentByID(c *gin.Context)
	GetCommentsByTarget(c *gin.Context)
	GetReplies(c *gin.Context)
	GetCommentsByUser(c *gin.Context)
}
