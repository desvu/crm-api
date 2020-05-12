package storefront

import (
	"context"
	"sort"

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

	if front.Blocks == nil {
		front.Blocks = []entity.Block{}
	}

	if err := s.verifyBlocks(ctx, front); err != nil {
		return nil, err
	}

	return front, s.Repository.Create(ctx, front)
}

func (s *Service) verifyBlocks(ctx context.Context, data *entity.Storefront) error {
	for i := range data.Blocks {
		if err := s.verifyBlock(ctx, &data.Blocks[i]); err != nil {
			return pkgerrors.Wrapf(err, "invalid block %d", i)
		}
	}
	return nil
}

func (s *Service) verifyBlock(ctx context.Context, data *entity.Block) error {
	if err := data.Validate(); err != nil {
		return err
	}
	err := s.GameRevisionService.IsGamesPublished(ctx, data.GameIDs...)
	if err != nil {
		return err
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

	if front.Blocks == nil {
		front.Blocks = []entity.Block{}
	}

	if err := s.verifyBlocks(ctx, front); err != nil {
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
	return s.Repository.FindByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*entity.Storefront, error) {
	result, err := s.Repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	sort.Slice(result, func(i, j int) bool {
		// active first
		if result[i].IsActive {
			return true
		}
		if result[j].IsActive {
			return false
		}
		// by last update
		return result[i].UpdatedAt.After(result[j].UpdatedAt)
	})
	return result, nil
}

func (s *Service) GetActive(ctx context.Context) (*entity.Storefront, error) {
	sf, err := s.Repository.FindActive(ctx)
	if err != nil {
		if err == errors.StoreFrontNotFound {
			return &entity.Storefront{
				IsActive: true,
				Blocks:   []entity.Block{},
			}, nil
		}
		return nil, err
	}
	return sf, nil
}
