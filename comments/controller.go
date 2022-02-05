package comments

import (
	"net/http"
	"strconv"

	"github.com/asrulkadir/section-comment-backend/pkg/validator"

	"github.com/labstack/echo/v4"
)

type commentController struct {
	service *ServiceComment
	valid   *validator.Validator
}

func NewController(e *echo.Echo, service *ServiceComment, valid *validator.Validator) {
	c := &commentController{service, valid}

	i := e.Group("/v1")
	i.GET("/data", c.GetData)
	i.GET("/comments", c.GetComments)
	// i.GET("/replies", c.GetReplies)
	i.POST("/comments", c.PostComment)
	i.POST("/replies", c.PostReply)
	i.PUT("/comments/:id", c.EditCommentContent)
	i.PUT("/replies/:id", c.EditReplyContent)
	i.PUT("/comments/:id/score", c.EditCommentScore)
	i.PUT("/replies/:id/score", c.EditReplyScore)
	i.DELETE("/comments/:id", c.DeleteComment)
	i.DELETE("/replies/:id", c.DeleteReply)
}

func (c *commentController) GetData(e echo.Context) error {
	currentUser, err := c.service.GetCurrentUser(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	comments, err := c.service.GetComments(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"current_user": currentUser,
		"comments":     comments,
	})
}

func (c *commentController) GetComments(e echo.Context) error {
	commentsData, err := c.service.GetComments(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"data":    commentsData,
		"message": "get comments success",
	})
}

func (c *commentController) PostComment(e echo.Context) error {
	u := new(CtrPostComment)

	if err := e.Bind(u); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := e.Validate(u); err != nil {
		return err
	}

	data := CtrPostComment{
		Content:  u.Content,
		Score:    u.Score,
		Username: u.Username,
	}

	err := c.service.PostComment(e.Request().Context(), data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "post comment success",
	})
}

func (c *commentController) DeleteComment(e echo.Context) error {
	id := e.Param("id")

	idComment, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = c.service.DeleteComment(e.Request().Context(), idComment)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete comment success",
	})
}

func (c *commentController) DeleteReply(e echo.Context) error {
	id := e.Param("id")

	idReply, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = c.service.DeleteReply(e.Request().Context(), idReply)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete reply success",
	})
}

func (c *commentController) PostReply(e echo.Context) error {
	u := new(CtrPostReply)

	if err := e.Bind(u); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := e.Validate(u); err != nil {
		return err
	}

	data := CtrPostReply{
		Content:    u.Content,
		Score:      u.Score,
		Username:   u.Username,
		ReplyingTo: u.ReplyingTo,
		IdComment:  u.IdComment,
	}

	err := c.service.PostReply(e.Request().Context(), data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "post reply success",
	})
}

func (c *commentController) EditCommentContent(e echo.Context) error {
	id := e.Param("id")

	idComment, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	u := new(CtrEditComment)

	if err := e.Bind(u); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := e.Validate(u); err != nil {
		return err
	}

	data := CtrEditComment{
		Content: u.Content,
	}

	err = c.service.EditCommentContent(e.Request().Context(), idComment, data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "edit comment content success",
	})
}

func (c *commentController) EditReplyContent(e echo.Context) error {
	id := e.Param("id")

	idReply, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	u := new(CtrEditReply)

	if err := e.Bind(u); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := e.Validate(u); err != nil {
		return err
	}

	data := CtrEditReply{
		Content: u.Content,
	}

	err = c.service.EditReplyContent(e.Request().Context(), idReply, data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "edit reply content success",
	})
}

func (c *commentController) EditCommentScore(e echo.Context) error {
	id := e.Param("id")

	idComment, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	u := new(CtrEditScore)

	if err := e.Bind(u); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := e.Validate(u); err != nil {
		return err
	}

	data := CtrEditScore{
		Score: u.Score,
	}

	err = c.service.EditCommentScore(e.Request().Context(), idComment, data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "edit comment score success",
	})
}

func (c *commentController) EditReplyScore(e echo.Context) error {
	id := e.Param("id")

	idReply, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	u := new(CtrEditScore)

	if err := e.Bind(u); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := e.Validate(u); err != nil {
		return err
	}

	data := CtrEditScore{
		Score: u.Score,
	}

	err = c.service.EditReplyScore(e.Request().Context(), idReply, data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "edit reply score success",
	})
}
