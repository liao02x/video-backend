package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info"  json:"info"  binding:"max=233"`
}

// Create video
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}

	if err := model.DB.Create(&video).Error; err == nil {
		return serializer.Response{
			Data: serializer.BuildVideo(video),
		}
	} else {
		return serializer.Response{
			Code:  50001,
			Msg:   "Created unsuccessfully",
			Error: err.Error(),
		}
	}

}
