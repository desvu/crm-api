package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameRevisionService interface {
	Update(ctx context.Context, data *UpdateGameRevisionData) (*entity.GameRevisionEx, error)

	GetByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error)
	GetExistByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error)
	GetDraftByGame(ctx context.Context, game *entity.Game) (*entity.GameRevisionEx, error)
}

type UpdateGameRevisionData struct {
	ID         uint
	Tags       *[]uint
	Developers *[]uint
	Publishers *[]uint
	Features   *[]uint
	Genres     *[]uint
}
