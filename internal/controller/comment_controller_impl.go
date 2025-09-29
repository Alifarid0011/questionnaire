package controller

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentControllerImpl struct {
	service service.CommentService
}

func NewCommentController(s service.CommentService) *CommentControllerImpl {
	return &CommentControllerImpl{service: s}
}

// CreateComment godoc
// @Summary Create a new comment
// @Description Adds a new comment to a target
// @Tags Comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /comments [post]
func (cc *CommentControllerImpl) CreateComment(c *gin.Context) {
	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	comment, err := cc.service.CreateComment(c.Request.Context(), &input)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(comment).Message("Comment created successfully").Dispatch()
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Updates an existing comment by ID
// @Tags Comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Updated comment data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /comments [put]
func (cc *CommentControllerImpl) UpdateComment(c *gin.Context) {
	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	comment, err := cc.service.UpdateComment(c.Request.Context(), &input)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(comment).Message("Comment updated successfully").Dispatch()
}

// GetCommentByID godoc
// @Summary Get comment by ID
// @Description Returns a comment by its ID
// @Tags Comments
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /comments/{id} [get]
func (cc *CommentControllerImpl) GetCommentByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	comment, err := cc.service.GetCommentByID(c.Request.Context(), id)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(comment).Dispatch()
}

// GetCommentsByTarget godoc
// @Summary Get comments by target
// @Description Returns all comments for a polymorphic target
// @Tags Comments
// @Produce json
// @Param ref query string true "Target collection"
// @Param id query string true "Target document ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /comments/target [get]
func (cc *CommentControllerImpl) GetCommentsByTarget(c *gin.Context) {
	ref := c.Query("ref")
	idParam := c.Query("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	target := models.DBRef{
		Ref: ref,
		ID:  id,
	}
	comments, err := cc.service.GetCommentsByTarget(c.Request.Context(), target)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(comments).Dispatch()
}

// GetReplies godoc
// @Summary Get replies for a comment
// @Description Returns all replies for a given parent comment
// @Tags Comments
// @Produce json
// @Param parent_id path string true "Parent comment ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /comments/{parent_id}/replies [get]
func (cc *CommentControllerImpl) GetReplies(c *gin.Context) {
	parentIDParam := c.Param("parent_id")
	parentID, err := primitive.ObjectIDFromHex(parentIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	replies, err := cc.service.GetReplies(c.Request.Context(), parentID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(replies).Dispatch()
}

// GetCommentsByUser godoc
// @Summary Get comments by user
// @Description Returns all comments created by a specific user
// @Tags Comments
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /comments/user/{user_id} [get]
func (cc *CommentControllerImpl) GetCommentsByUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	comments, err := cc.service.GetCommentsByUser(c.Request.Context(), userID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(comments).Dispatch()
}
