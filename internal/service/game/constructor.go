package game

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type ServiceParams struct {
	TagService     service.TagService
	GameRepository repository.IGameRepository
}

func New(params ServiceParams) *Service {
	return &Service{
		params,
	}
}
