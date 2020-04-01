package game

import (
	"github.com/qilin/crm-api/internal/domain/repository"
)

type ServiceParams struct {
	GameRepository repository.IGameRepository
}

type Service struct {
	ServiceParams
}

func New(params ServiceParams) *Service {
	return &Service{
		params,
	}
}
