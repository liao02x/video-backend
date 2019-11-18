package api

import (
	//"singo/serializer"
	"singo/service"

	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/sessions"
)

// CreateVideo
func CreateVideo(c *gin.Context) {
	s := service.CreateVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo
func ShowVideo(c *gin.Context) {
	s := service.ShowVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListVideo
func ListVideo(c *gin.Context) {
	s := service.ListVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateVideo
func UpdateVideo(c *gin.Context) {
	s := service.UpdateVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo
func DeleteVideo(c *gin.Context) {
	s := service.DeleteVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}