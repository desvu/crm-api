package game

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type ServiceParams struct {
	GameRepository repository.IGameRepository
}

func New(params ServiceParams) service.IGameService {
	return &Service{
		params,
	}
}
