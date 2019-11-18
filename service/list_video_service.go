package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService
type ListVideoService struct {
}

// Show
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video

	if err := model.DB.Find(&videos).Error; err == nil {
		return serializer.Response{
			Data: serializer.BuildVideos(videos),
		}
	} else {
		return serializer.Response{
			Code:  50000,
			Msg:   "Unknown server or database error",
			Error: err.Error(),
		}
	}

}
