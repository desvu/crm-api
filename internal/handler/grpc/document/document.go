package document

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/errors/grpcerror"
	"github.com/qilin/crm-api/pkg/grpc/proto"
)

func (h *Handler) AddDocument(ctx context.Context, request *proto.AddDocumentRequest) (*proto.DocumentResponse, error) {
	doc, err := h.DocumentService.GetByID(ctx, uint(request.DocID))
	if err != nil {
		return nil, err
	}

	user, err := h.UserService.GetByExternalID(ctx, request.UserID)
	if err != nil {
		return nil, err
	}

	err = h.DocumentService.AddDocumentToUser(ctx, doc, user.ID)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	result, err := h.convertDocument(doc)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	return &proto.DocumentResponse{Document: result}, nil
}

func (h *Handler) GetByID(ctx context.Context, request *proto.GetDocumentByIDRequest) (*proto.DocumentResponse, error) {
	doc, err := h.DocumentService.GetByID(ctx, uint(request.DocumentID))
	if err != nil {
		return nil, grpcerror.New(err)
	}

	// return only activated document
	if doc.ActivatedAt == nil {
		return nil, grpcerror.New(errors.DocumentNotFound)
	}

	result, err := h.convertDocument(doc)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	return &proto.DocumentResponse{Document: result}, nil
}

func (h *Handler) GetByUserID(ctx context.Context, request *proto.GetDocumentsByUserIDRequest) (*proto.DocumentsResponse, error) {
	if request.Limit == 0 {
		request.Limit = 30
	}

	user, err := h.UserService.GetByExternalID(ctx, request.UserID)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	docs, err := h.DocumentService.GetUserDocuments(ctx, &service.GetUserDocumentsData{
		UserID: user.ID,
		Limit:  uint(request.Limit),
		Offset: uint(request.Offset),
	})

	result, err := h.convertDocuments(docs)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	return &proto.DocumentsResponse{Documents: result}, nil
}

func (h *Handler) convertDocument(doc *entity.Document) (*proto.Document, error) {
	return &proto.Document{
		ID:       uint64(doc.ID),
		Title:    doc.Title,
		Text:     doc.Text,
		Type:     doc.Type.String(),
		Language: string(doc.Language),
		Version:  doc.Version,
	}, nil
}

func (h *Handler) convertDocuments(docs []entity.Document) ([]*proto.Document, error) {
	var result []*proto.Document
	for _, item := range docs {
		g, err := h.convertDocument(&item)
		if err != nil {
			return nil, err
		}

		result = append(result, g)
	}

	return result, nil
}
