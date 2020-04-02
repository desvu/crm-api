package service

import (
	"github.com/qilin/crm-api/internal/app/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/internal/service/game"
	"github.com/qilin/crm-api/internal/service/tag"
)

type Services struct {
	GameService service.IGameService
	TagService  service.ITagService
}

func New(r *repository.Repositories) *Services {
	s := new(Services)

	s.GameService = game.New(
		game.ServiceParams{
			GameRepository: r.GameRepository,
		},
	)

	s.TagService = tag.New(
		tag.ServiceParams{
			GameService:       s.GameService,
			TagRepository:     r.TagRepository,
			GameTagRepository: r.GameTagRepository,
		},
	)

	return s
}
