package localization

import (
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
	"go.uber.org/fx"
)

type ServiceParams struct {
	fx.In

	LocalizationRepository repository.GameRevisionLocalizationRepository
}

func New(params ServiceParams) service.LocalizationService {
	return &Service{
		params,
	}
}
