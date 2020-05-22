package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/review_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service ReviewService
type ReviewService interface {
	//GetByID(ctx context.Context, id uint) (*entity.Review, error)
	//GetExistByID(ctx context.Context, id uint) (*entity.Review, error)
	//GetByIDs(ctx context.Context, ids []uint) ([]entity.Review, error)
	GetByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Review, error)

	UpdateReviewsForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, reviews []ReviewData) error
}

type ReviewData struct {
	PressName string `validate:"required"`
	Link      string `validate:"required"`
	Score     uint   `validate:"required,lte=100"`
	Quote     string `validate:"required"`
}

func (d ReviewData) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
