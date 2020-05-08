package storefront

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/pkg/errors/grpcerror"
	"github.com/qilin/crm-api/pkg/grpc/proto"
)

func (h *Handler) GetActive(ctx context.Context, request *proto.GetActiveStorefrontRequest) (*proto.StorefrontResponse, error) {
	storefront, err := h.StorefrontService.GetActive(ctx)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	result, err := h.convert(storefront)
	if err != nil {
		return nil, grpcerror.New(err)
	}

	return &proto.StorefrontResponse{Storefront: result}, nil
}

func (h *Handler) convert(sf *entity.Storefront) (*proto.Storefront, error) {
	res := &proto.Storefront{}

	for _, b := range sf.Blocks {
		res.Blocks = append(res.Blocks, &proto.Block{
			Type:    b.Type.String(),
			Title:   b.Title.String(),
			Filter:  b.Filter,
			GameIds: b.GameIDs,
		})
	}

	return res, nil
}
