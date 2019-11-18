package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService
type ShowVideoService struct {
}

// Show
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video

	if err := model.DB.First(&video, id).Error; err == nil {
		return serializer.Response{
			Data: serializer.BuildVideo(video),
		}
	} else {
		return serializer.Response{
			Code:  404,
			Msg:   "Video not found",
			Error: err.Error(),
		}
	}

}
