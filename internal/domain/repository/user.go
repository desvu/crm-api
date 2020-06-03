package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/user_document_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository UserDocumentRepository
type UserRepository interface {
	FindByExternalID(ctx context.Context, externalID string) (*entity.User, error)
}
