package app

import "service_social_media/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Post{}},
		{Model: models.PostImage{}},
		{Model: models.Comment{}},
		{Model: models.Like{}},
	}
}
