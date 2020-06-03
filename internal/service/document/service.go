package document

import (
	"context"
	"fmt"
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s *Service) Create(ctx context.Context, data *service.CommonDocumentData) (*entity.Document, error) {
	// todo: validate data
	doc := &entity.Document{
		Title:       data.Title,
		Text:        data.Text,
		Reason:      data.Reason,
		Type:        data.Type,
		Language:    data.Language,
		Version:     data.Version,
		CreatedBy:   data.CreatedBy,
		ActivatedBy: "",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		ActivatedAt: nil,
	}
	if err := s.DocumentRepository.Create(ctx, doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *Service) Update(ctx context.Context, data *service.UpdateDocumentData) (*entity.Document, error) {
	doc, err := s.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if doc.ActivatedAt != nil {
		return nil, errors.DocumentAlreadyActivated
	}

	doc.Reason = data.Reason
	doc.Text = data.Text
	doc.Type = data.Type
	now := time.Now()
	doc.UpdatedAt = &now

	err = s.DocumentRepository.Update(ctx, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *Service) Upsert(ctx context.Context, data *service.UpsertDocumentData) (*entity.Document, error) {
	if data.ID != nil {
		return s.Update(ctx, &service.UpdateDocumentData{
			ID: *data.ID,
			CommonDocumentData: service.CommonDocumentData{
				Text:     data.Text,
				Reason:   data.Reason,
				Type:     data.Type,
				Language: data.Language,
				Version:  data.Version,
			},
		})
	}
	return s.Create(ctx, &service.CommonDocumentData{
		Text:     data.Text,
		Reason:   data.Reason,
		Type:     data.Type,
		Language: data.Language,
		Version:  data.Version,
	})
}

func (s *Service) Delete(ctx context.Context, id uint) error {
	doc, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.DocumentRepository.Delete(ctx, doc)
}

func (s *Service) Activate(ctx context.Context, id uint) error {
	doc, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	now := time.Now()
	loc, ok := titleNames[doc.Language]
	if !ok {
		return errors.DocumentUnsupportedLanguage
	}
	title, ok := loc[doc.Type]
	if !ok {
		return errors.DocumentUnsupportedLanguage
	}

	doc.Title = fmt.Sprintf("%s %s", title, doc.Version)
	doc.ActivatedAt = &now
	//doc.ActivatedBy = "" // todo

	err = s.DocumentRepository.Update(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddDocumentToUser(ctx context.Context, document *entity.Document, userID uint) error {
	if document.ActivatedAt == nil {
		return errors.DocumentNotFound
	}

	userDoc, err := s.UserDocumentRepository.FindByUserAndDocumentID(ctx, userID, document.ID)
	if err != nil {
		return err
	}
	if userDoc != nil {
		return errors.DocumentAlreadyAdded
	}

	return s.UserDocumentRepository.Create(ctx, &entity.UserDocument{
		UserID:     userID,
		DocumentID: document.ID,
		CreatedAt:  time.Now(),
	})
}

func (s *Service) GetByID(ctx context.Context, id uint) (*entity.Document, error) {
	doc, err := s.DocumentRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if doc == nil {
		return nil, errors.DocumentNotFound
	}

	return doc, nil
}

func (s *Service) GetExByID(ctx context.Context, id uint) (*entity.Document, error) {
	return s.DocumentRepository.FindByID(ctx, id)
}

// GetExByFilter returns games founded by filter
func (s *Service) GetExByFilter(ctx context.Context, data *service.GetByFilterDocumentData) ([]entity.Document, error) {
	if data.Limit == 0 {
		data.Limit = 30
	}

	return s.DocumentRepository.FindByFilter(ctx, &repository.FindByFilterDocumentData{
		OnlyActivated: data.OnlyActivated,
		Limit:         data.Limit,
		Offset:        data.Offset,
	})
}

func (s *Service) GetCountByFilter(ctx context.Context, data *service.GetByFilterDocumentData) (int, error) {
	return s.DocumentRepository.CountByFilter(ctx, &repository.FindByFilterDocumentData{
		OnlyActivated: data.OnlyActivated,
		Limit:         data.Limit,
		Offset:        data.Offset,
	})
}

func (s *Service) GetUserDocuments(ctx context.Context, data *service.GetUserDocumentsData) ([]entity.Document, error) {
	if data.Limit == 0 {
		data.Limit = 10
	}
	if data.Limit > 1000 {
		data.Limit = 1000
	}
	userDocs, err := s.UserDocumentRepository.FindByUserID(ctx, repository.FindUserDocumentsByUserIdData{
		UserID: data.UserID,
		Limit:  int(data.Limit),
		Offset: int(data.Offset),
	})
	if err != nil {
		return nil, err
	}

	return s.DocumentRepository.FindByIDs(ctx, entity.NewUserDocumentArray(userDocs).DocumentIDs())
}
