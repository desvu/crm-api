package graph

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/internal/handler/graph/model"
)

// CreateFeature creates new feature in db
func (r *mutationResolver) CreateFeature(ctx context.Context, name string, icon string) (*model.Feature, error) {
	feature, err := r.featureService.Create(ctx, &service.CreateFeatureData{
		Name: name,
		Icon: game.NewIcon(icon),
	})
	if err != nil {
		return nil, err
	}
	return r.convertFeature(*feature), nil
}
