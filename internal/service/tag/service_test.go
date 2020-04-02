package tag

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/mocks"
)

func TestService_sanitizeAttachTags(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	gameService := mocks.NewMockIGameService(ctrl)
	tagRepository := mocks.NewMockITagRepository(ctrl)
	gameTagRepository := mocks.NewMockIGameTagRepository(ctrl)

	s := New(ServiceParams{
		GameService:       gameService,
		TagRepository:     tagRepository,
		GameTagRepository: gameTagRepository,
	})

	type args struct {
		ctx    context.Context
		gameID uint
		tagIDs []uint
	}
	tests := []struct {
		name    string
		args    args
		want    []uint
		wantErr bool
	}{
		{
			name: "getting positive data",
			args: args{
				ctx:    context.Background(),
				gameID: 1,
				tagIDs: []uint{1, 2, 3},
			},
			want:    []uint{1, 2},
			wantErr: false,
		},
		{
			name: "getting a non-existent tag ID",
			args: args{
				ctx:    context.Background(),
				gameID: 1,
				tagIDs: []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "getting a non-existent game ID",
			args: args{
				ctx:    context.Background(),
				gameID: 2,
				tagIDs: []uint{1, 2, 3},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// s.GameService.GetExistByID
			gameService.EXPECT().GetExistByID(gomock.Any(), gomock.Any()).Return(&entity.Game{ID: 1}, nil)

			// s.GetByIDs
			tagRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).Return([]entity.Tag{{ID: 1}, {ID: 2}, {ID: 3}}, nil).Times(1)

			// s.GetByGameID
			gameTagRepository.EXPECT().FindByGameID(gomock.Any(), gomock.Any()).Return([]entity.GameTag{{TagID: 3}}, nil)
			tagRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).Return([]entity.Tag{{ID: 3}}, nil)

			got, err := s.sanitizeAttachTags(tt.args.ctx, tt.args.gameID, tt.args.tagIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("sanitizeAttachTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sanitizeAttachTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_sanitizeDetachTags(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	gameService := mocks.NewMockIGameService(ctrl)
	tagRepository := mocks.NewMockITagRepository(ctrl)
	gameTagRepository := mocks.NewMockIGameTagRepository(ctrl)

	s := New(ServiceParams{
		GameService:       gameService,
		TagRepository:     tagRepository,
		GameTagRepository: gameTagRepository,
	})

	type args struct {
		ctx    context.Context
		gameID uint
		tagIDs []uint
	}
	tests := []struct {
		name    string
		args    args
		want    []uint
		wantErr bool
	}{
		{
			name: "getting positive data",
			args: args{
				ctx:    context.Background(),
				gameID: 1,
				tagIDs: []uint{1, 2, 3},
			},
			want:    []uint{3},
			wantErr: false,
		},
		{
			name: "getting a non-existent tag ID",
			args: args{
				ctx:    context.Background(),
				gameID: 1,
				tagIDs: []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
		{
			name: "getting a non-existent game ID",
			args: args{
				ctx:    context.Background(),
				gameID: 2,
				tagIDs: []uint{1, 2, 3},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// s.GameService.GetExistByID
			gameService.EXPECT().GetExistByID(gomock.Any(), gomock.Any()).Return(&entity.Game{ID: 1}, nil)

			// s.GetByIDs
			tagRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).Return([]entity.Tag{{ID: 1}, {ID: 2}, {ID: 3}}, nil).Times(1)

			// s.GetByGameID
			gameTagRepository.EXPECT().FindByGameID(gomock.Any(), gomock.Any()).Return([]entity.GameTag{{TagID: 3}}, nil)
			tagRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).Return([]entity.Tag{{ID: 3}}, nil)

			got, err := s.sanitizeDetachTags(tt.args.ctx, tt.args.gameID, tt.args.tagIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("sanitizeDetachTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sanitizeDetachTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}
