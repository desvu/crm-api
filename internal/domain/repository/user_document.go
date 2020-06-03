package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/user_document_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository UserDocumentRepository
type UserDocumentRepository interface {
	Create(ctx context.Context, i *entity.UserDocument) error

	FindByUserID(ctx context.Context, data FindUserDocumentsByUserIdData) ([]entity.UserDocument, error)
	FindByUserAndDocumentID(ctx context.Context, userId, docId uint) (*entity.UserDocument, error)
}

type FindUserDocumentsByUserIdData struct {
	UserID uint
	Limit  int
	Offset int
}
