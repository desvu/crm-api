package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID         uint      `pg:"id"`
	UserID     uint      `pg:"user_id"`
	DocumentID uint      `pg:"document_id"`
	CreatedAt  time.Time `pg:"created_at"`

	tableName struct{} `pg:"user_documents"`
}

func (m model) Convert() *entity.UserDocument {
	return &entity.UserDocument{
		ID:         m.ID,
		UserID:     m.UserID,
		DocumentID: m.DocumentID,
		CreatedAt:  m.CreatedAt,
	}
}

func newModel(i *entity.UserDocument) (*model, error) {
	return &model{
		ID:         i.ID,
		UserID:     i.UserID,
		DocumentID: i.DocumentID,
		CreatedAt:  i.CreatedAt,
	}, nil
}
