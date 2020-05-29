package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/document"
	"github.com/qilin/crm-api/internal/domain/enum/language"
)

type model struct {
	ID          uint              `pg:"id"`
	Title       string            `pg:"title"`
	Text        string            `pg:"text"`
	Reason      string            `pg:"reason"`
	Type        document.Type     `pg:"type"`
	Language    language.Language `pg:"language"`
	Version     string            `pg:"version""`
	CreatedBy   string            `pg:"created_by"`
	ActivatedBy string            `pg:"activated_by"`
	CreatedAt   time.Time         `pg:"created_at"`
	UpdatedAt   *time.Time        `pg:"updated_at"`
	ActivatedAt *time.Time        `pg:"activated_at"`

	tableName struct{} `pg:"documents"`
}

func (m model) Convert() *entity.Document {
	return &entity.Document{
		ID:          m.ID,
		Title:       m.Title,
		Text:        m.Text,
		Reason:      m.Reason,
		Type:        m.Type,
		Language:    m.Language,
		Version:     m.Version,
		CreatedBy:   m.CreatedBy,
		ActivatedBy: m.ActivatedBy,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		ActivatedAt: m.ActivatedAt,
	}
}

func newModel(i *entity.Document) (*model, error) {
	return &model{
		ID:          i.ID,
		Title:       i.Title,
		Text:        i.Text,
		Reason:      i.Reason,
		Type:        i.Type,
		Language:    i.Language,
		Version:     i.Version,
		CreatedBy:   i.CreatedBy,
		ActivatedBy: i.ActivatedBy,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		ActivatedAt: i.ActivatedAt,
	}, nil
}
