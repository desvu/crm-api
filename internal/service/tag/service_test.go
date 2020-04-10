package tag

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/mocks"
)

func TestService_getGameTagsForInsert(t *testing.T) {
	t.Parallel()

	type args struct {
		gameID          uint
		newTagIDs       []uint
		currentGameTags []entity.GameTag
	}
	tests := []struct {
		name string
		args args
		want []entity.GameTag
	}{
		{
			name: "getting a list of tag IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:          1,
				newTagIDs:       []uint{1, 2, 3, 4},
				currentGameTags: []entity.GameTag{{TagID: 2}, {TagID: 3}},
			},
			want: []entity.GameTag{{TagID: 1, GameID: 1}, {TagID: 4, GameID: 1}},
		},
		{
			name: "getting a list of tag IDs with a fully included subset of IDs associated with the device",
			args: args{
				gameID:          1,
				newTagIDs:       []uint{2, 3},
				currentGameTags: []entity.GameTag{{TagID: 2}, {TagID: 3}},
			},
			want: []entity.GameTag{},
		},
		{
			name: "getting a list of tag IDs with or without an incoming subset of IDs associated with the device",
			args: args{
				gameID:          1,
				newTagIDs:       []uint{5, 6},
				currentGameTags: []entity.GameTag{{TagID: 2}, {TagID: 3}},
			},
			want: []entity.GameTag{{TagID: 5, GameID: 1}, {TagID: 6, GameID: 1}},
		},
		{
			name: "getting a list of tag IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:          1,
				newTagIDs:       []uint{1, 2, 2, 3, 4},
				currentGameTags: []entity.GameTag{{TagID: 2}, {TagID: 3}},
			},
			want: []entity.GameTag{{TagID: 1, GameID: 1}, {TagID: 4, GameID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameTagsForInsert(tt.args.gameID, tt.args.newTagIDs, tt.args.currentGameTags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameTagsForInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_getGameTagsForDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		newTagIDs       []uint
		currentGameTags []entity.GameTag
	}
	tests := []struct {
		name string
		args args
		want []entity.GameTag
	}{
		{
			name: "getting a list of tag IDs with a partially included subset of IDs associated with the device",
			args: args{
				newTagIDs:       []uint{1, 2, 3, 4},
				currentGameTags: []entity.GameTag{{TagID: 2}, {TagID: 3}},
			},
			want: []entity.GameTag{},
		},
		{
			name: "getting a list of tag IDs with a fully included subset of IDs associated with the device",
			args: args{
				newTagIDs:       []uint{2, 3},
				currentGameTags: []entity.GameTag{{TagID: 2}, {TagID: 3}},
			},
			want: []entity.GameTag{},
		},
		{
			name: "getting a list of tag IDs with a partially included subset of IDs associated with the device",
			args: args{
				newTagIDs:       []uint{1, 4},
				currentGameTags: []entity.GameTag{{ID: 1, TagID: 2, GameID: 1}, {ID: 1, TagID: 3, GameID: 1}},
			},
			want: []entity.GameTag{{ID: 1, TagID: 2, GameID: 1}, {ID: 1, TagID: 3, GameID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameTagsForDelete(tt.args.newTagIDs, tt.args.currentGameTags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameTagsForDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateTagsForGame(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	tagRepository := mocks.NewMockTagRepository(ctrl)
	gameTagRepository := mocks.NewMockGameTagRepository(ctrl)

	s := New(ServiceParams{
		TagRepository:     tagRepository,
		GameTagRepository: gameTagRepository,
	})

	type args struct {
		ctx    context.Context
		game   *entity.Game
		tagIDs []uint
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Tag
		wantErr bool
	}{
		{
			name: "getting a non-existent tag ID",
			args: args{
				ctx:    context.Background(),
				game:   &entity.Game{ID: 1},
				tagIDs: []uint{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "getting a non-existent tag ID",
			args: args{
				ctx:    context.Background(),
				game:   &entity.Game{ID: 1},
				tagIDs: []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s.GetByIDs
			tagRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).
				Return([]entity.Tag{{ID: 1}, {ID: 2}, {ID: 3}}, nil)

			// s.GameTagRepository.FindByGameID
			gameTagRepository.EXPECT().FindByGameID(gomock.Any(), gomock.Any()).
				Return([]entity.GameTag{{TagID: 3}}, nil)

			gameTagRepository.EXPECT().DeleteMultiple(gomock.Any(), gomock.Any())
			gameTagRepository.EXPECT().CreateMultiple(gomock.Any(), gomock.Any())

			err := s.UpdateTagsForGame(tt.args.ctx, tt.args.game, tt.args.tagIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTagsForGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
