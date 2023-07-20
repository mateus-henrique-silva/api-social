package core

import (
	"context"

	"go.mod/src/db"
	"go.mod/src/entity"
	"go.mod/src/rest"
)

type youtubeManager struct {
}

func NewYoutubeManager() *youtubeManager {
	return &youtubeManager{}
}

func (m *youtubeManager) GetYoutube(ctx context.Context) ([]entity.Youtube, error) {
	result, err := db.ReturnAllVideos(ctx)
	if err != nil {
		return nil, &rest.Error{Status: 400, Code: "erro_consult", Message: err.Error()}
	}
	return result, nil

}

func (m *youtubeManager) PostYoutube(ctx context.Context, youtube entity.Youtube) error {
	err := db.CreateYoutube(ctx, youtube)
	if err != nil {
		return &rest.Error{Status: 400, Code: "erro_post_youtube", Message: "erro ao criar post"}
	}
	return nil
}

func (m *youtubeManager) PutYoutube(ctx context.Context, youtube entity.Youtube, id string) error {
	err := db.PutYotube(ctx, youtube, id)
	if err != nil {
		return &rest.Error{Status: 400, Code: "erro_put_youtube.", Message: "erro ao atualizar o video"}
	}
	return nil
}
