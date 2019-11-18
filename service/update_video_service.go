package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info"  json:"info"  binding:"required,min=0,max=233"`
}

// Show
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video

	if err := model.DB.First(&video, id).Error; err == nil {

	} else {
		return serializer.Response{
			Code:  404,
			Msg:   "Video not found",
			Error: err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info

	if err := model.DB.Save(&video).Error; err == nil {
		return serializer.Response{
			Data: serializer.BuildVideo(video),
		}
	} else {
		return serializer.Response{
			Code:  50002,
			Msg:   "Video updated unsuccessfully",
			Error: err.Error(),
		}
	}

}
