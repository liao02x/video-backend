package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteVideoService
type DeleteVideoService struct {
}

// Delete
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video

	if err := model.DB.First(&video, id).Error; err == nil {

	} else {
		return serializer.Response{
			Code:  404,
			Msg:   "Video not found",
			Error: err.Error(),
		}
	}

	if err := model.DB.Delete(&video).Error; err == nil {
		return serializer.Response{}
	} else {
		return serializer.Response{
			Code:  50003,
			Msg:   "Video deleted unsuccessfully",
			Error: err.Error(),
		}
	}

}
