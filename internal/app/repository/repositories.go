package repository

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/repository/game"
	"github.com/qilin/crm-api/internal/repository/game_tag"
	"github.com/qilin/crm-api/internal/repository/tag"
)

type Repositories struct {
	GameRepository    repository.IGameRepository
	TagRepository     repository.TagRepository
	GameTagRepository repository.GameTagRepository
}

func New(e *env.Store) *Repositories {
	return &Repositories{
		GameRepository:    game.New(e),
		TagRepository:     tag.New(e),
		GameTagRepository: game_tag.New(e),
	}
}
