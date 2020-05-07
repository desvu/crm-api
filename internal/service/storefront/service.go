package storefront

import (
	"context"

	pkgerrors "github.com/pkg/errors"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/errors"
)

type Service struct {
	ServiceParams
}

func (s *Service) Create(ctx context.Context, data *entity.Storefront) (*entity.Storefront, error) {
	front := &entity.Storefront{
		Name:     data.Name,
		Blocks:   data.Blocks,
		Version:  0,
		IsActive: false,
	}

	if err := s.verifyGames(ctx, front); err != nil {
		return nil, err
	}

	return front, s.Repository.Create(ctx, front)
}

func (s *Service) verifyGames(ctx context.Context, data *entity.Storefront) error {
	for i := range data.Blocks {
		if err := s.GameRevisionService.IsGamesPublished(ctx, data.Blocks[i].GameIDs...); err != nil {
			return pkgerrors.Wrapf(err, "invalid games for block %d", i)
		}
	}
	return nil
}

func (s *Service) Update(ctx context.Context, data *entity.Storefront) (*entity.Storefront, error) {
	front, err := s.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	front.Name = data.Name
	front.Blocks = data.Blocks
	front.Version++

	if err := s.verifyGames(ctx, front); err != nil {
		return nil, err
	}

	return front, s.Repository.Update(ctx, front)
}

func (s *Service) Delete(ctx context.Context, id uint) error {
	front, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if front.IsActive {
		return errors.StoreFrontIsActive
	}

	return s.Repository.Delete(ctx, id)
}

func (s *Service) Activate(ctx context.Context, id uint) error {
	front, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.Repository.Activate(ctx, front.ID, front.Version)
}

func (s *Service) GetByID(ctx context.Context, id uint) (*entity.Storefront, error) {
	return s.Repository.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context, id uint) ([]*entity.Storefront, error) {
	return s.Repository.GetAll(ctx)
}

func (s *Service) FindActive(ctx context.Context) (*entity.Storefront, error) {
	return s.Repository.FindActive(ctx)
}