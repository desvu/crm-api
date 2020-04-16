package genre

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/mocks"
)

func TestService_getGameGenresForInsert(t *testing.T) {
	t.Parallel()

	type args struct {
		gameID            uint
		newGenreIDs       []uint
		currentGameGenres []entity.GameRevisionGenre
	}
	tests := []struct {
		name string
		args args
		want []entity.GameRevisionGenre
	}{
		{
			name: "getting a list of genre IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:            1,
				newGenreIDs:       []uint{1, 2, 3, 4},
				currentGameGenres: []entity.GameRevisionGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameRevisionGenre{{GenreID: 1, GameRevisionID: 1}, {GenreID: 4, GameRevisionID: 1}},
		},
		{
			name: "getting a list of genre IDs with a fully included subset of IDs associated with the device",
			args: args{
				gameID:            1,
				newGenreIDs:       []uint{2, 3},
				currentGameGenres: []entity.GameRevisionGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameRevisionGenre{},
		},
		{
			name: "getting a list of genre IDs with or without an incoming subset of IDs associated with the device",
			args: args{
				gameID:            1,
				newGenreIDs:       []uint{5, 6},
				currentGameGenres: []entity.GameRevisionGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameRevisionGenre{{GenreID: 5, GameRevisionID: 1}, {GenreID: 6, GameRevisionID: 1}},
		},
		{
			name: "getting a list of genre IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:            1,
				newGenreIDs:       []uint{1, 2, 2, 3, 4},
				currentGameGenres: []entity.GameRevisionGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameRevisionGenre{{GenreID: 1, GameRevisionID: 1}, {GenreID: 4, GameRevisionID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameGenresForInsert(tt.args.gameID, tt.args.newGenreIDs, tt.args.currentGameGenres); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameGenresForInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_getGameGenresForDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		newGenreIDs       []uint
		currentGameGenres []entity.GameRevisionGenre
	}
	tests := []struct {
		name string
		args args
		want []entity.GameRevisionGenre
	}{
		{
			name: "getting a list of genre IDs with a partially included subset of IDs associated with the device",
			args: args{
				newGenreIDs:       []uint{1, 2, 3, 4},
				currentGameGenres: []entity.GameRevisionGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameRevisionGenre{},
		},
		{
			name: "getting a list of genre IDs with a fully included subset of IDs associated with the device",
			args: args{
				newGenreIDs:       []uint{2, 3},
				currentGameGenres: []entity.GameRevisionGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameRevisionGenre{},
		},
		{
			name: "getting a list of genre IDs with a partially included subset of IDs associated with the device",
			args: args{
				newGenreIDs:       []uint{1, 4},
				currentGameGenres: []entity.GameRevisionGenre{{ID: 1, GenreID: 2, GameRevisionID: 1}, {ID: 1, GenreID: 3, GameRevisionID: 1}},
			},
			want: []entity.GameRevisionGenre{{ID: 1, GenreID: 2, GameRevisionID: 1}, {ID: 1, GenreID: 3, GameRevisionID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameGenresForDelete(tt.args.newGenreIDs, tt.args.currentGameGenres); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameGenresForDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateGenresForGame(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	genreRepository := mocks.NewMockGenreRepository(ctrl)
	gameRevisionGenreRepository := mocks.NewMockGameRevisionGenreRepository(ctrl)

	s := New(ServiceParams{
		GenreRepository:             genreRepository,
		GameRevisionGenreRepository: gameRevisionGenreRepository,
	})

	type args struct {
		ctx          context.Context
		gameRevision *entity.GameRevision
		genreIDs     []uint
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Genre
		wantErr bool
	}{
		{
			name: "getting a non-existent genre ID",
			args: args{
				ctx:          context.Background(),
				gameRevision: &entity.GameRevision{ID: 1},
				genreIDs:     []uint{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "getting a non-existent genre ID",
			args: args{
				ctx:          context.Background(),
				gameRevision: &entity.GameRevision{ID: 1},
				genreIDs:     []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s.GetByIDs
			genreRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).
				Return([]entity.Genre{{ID: 1}, {ID: 2}, {ID: 3}}, nil)

			// s.GameRevisionGenreRepository.FindByGameRevisionID
			gameRevisionGenreRepository.EXPECT().FindByGameRevisionID(gomock.Any(), gomock.Any()).
				Return([]entity.GameRevisionGenre{{GenreID: 3}}, nil)

			gameRevisionGenreRepository.EXPECT().DeleteMultiple(gomock.Any(), gomock.Any())
			gameRevisionGenreRepository.EXPECT().CreateMultiple(gomock.Any(), gomock.Any())

			err := s.UpdateGenresForGameRevision(tt.args.ctx, tt.args.gameRevision, tt.args.genreIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateGenresForGameRevision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
