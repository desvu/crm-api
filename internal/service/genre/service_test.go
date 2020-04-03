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

	ctrl := gomock.NewController(t)
	gameService := mocks.NewMockGameService(ctrl)
	genreRepository := mocks.NewMockGenreRepository(ctrl)
	gameGenreRepository := mocks.NewMockGameGenreRepository(ctrl)

	s := New(ServiceParams{
		GameService:         gameService,
		GenreRepository:     genreRepository,
		GameGenreRepository: gameGenreRepository,
	})

	type args struct {
		gameID            uint
		newGenreIDs       []uint
		currentGameGenres []entity.GameGenre
	}
	tests := []struct {
		name string
		args args
		want []entity.GameGenre
	}{
		{
			name: "getting a list of genre IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:            1,
				newGenreIDs:       []uint{1, 2, 3, 4},
				currentGameGenres: []entity.GameGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameGenre{{GenreID: 1, GameID: 1}, {GenreID: 4, GameID: 1}},
		},
		{
			name: "getting a list of genre IDs with a fully included subset of IDs associated with the device",
			args: args{
				gameID:            1,
				newGenreIDs:       []uint{2, 3},
				currentGameGenres: []entity.GameGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameGenre{},
		},
		{
			name: "getting a list of genre IDs with or without an incoming subset of IDs associated with the device",
			args: args{
				gameID:            1,
				newGenreIDs:       []uint{5, 6},
				currentGameGenres: []entity.GameGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameGenre{{GenreID: 5, GameID: 1}, {GenreID: 6, GameID: 1}},
		},
		{
			name: "getting a list of genre IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:            1,
				newGenreIDs:       []uint{1, 2, 2, 3, 4},
				currentGameGenres: []entity.GameGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameGenre{{GenreID: 1, GameID: 1}, {GenreID: 4, GameID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.getGameGenresForInsert(tt.args.gameID, tt.args.newGenreIDs, tt.args.currentGameGenres); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameGenresForInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_getGameGenresForDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	gameService := mocks.NewMockGameService(ctrl)
	genreRepository := mocks.NewMockGenreRepository(ctrl)
	gameGenreRepository := mocks.NewMockGameGenreRepository(ctrl)

	s := New(ServiceParams{
		GameService:         gameService,
		GenreRepository:     genreRepository,
		GameGenreRepository: gameGenreRepository,
	})

	type args struct {
		newGenreIDs       []uint
		currentGameGenres []entity.GameGenre
	}
	tests := []struct {
		name string
		args args
		want []entity.GameGenre
	}{
		{
			name: "getting a list of genre IDs with a partially included subset of IDs associated with the device",
			args: args{
				newGenreIDs:       []uint{1, 2, 3, 4},
				currentGameGenres: []entity.GameGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameGenre{},
		},
		{
			name: "getting a list of genre IDs with a fully included subset of IDs associated with the device",
			args: args{
				newGenreIDs:       []uint{2, 3},
				currentGameGenres: []entity.GameGenre{{GenreID: 2}, {GenreID: 3}},
			},
			want: []entity.GameGenre{},
		},
		{
			name: "getting a list of genre IDs with a partially included subset of IDs associated with the device",
			args: args{
				newGenreIDs:       []uint{1, 4},
				currentGameGenres: []entity.GameGenre{{ID: 1, GenreID: 2, GameID: 1}, {ID: 1, GenreID: 3, GameID: 1}},
			},
			want: []entity.GameGenre{{ID: 1, GenreID: 2, GameID: 1}, {ID: 1, GenreID: 3, GameID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.getGameGenresForDelete(tt.args.newGenreIDs, tt.args.currentGameGenres); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameGenresForDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateGenresForGame(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	gameService := mocks.NewMockGameService(ctrl)
	genreRepository := mocks.NewMockGenreRepository(ctrl)
	gameGenreRepository := mocks.NewMockGameGenreRepository(ctrl)

	s := New(ServiceParams{
		GameService:         gameService,
		GenreRepository:     genreRepository,
		GameGenreRepository: gameGenreRepository,
	})

	type args struct {
		ctx      context.Context
		game     *entity.Game
		genreIDs []uint
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
				ctx:      context.Background(),
				game:     &entity.Game{ID: 1},
				genreIDs: []uint{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "getting a non-existent genre ID",
			args: args{
				ctx:      context.Background(),
				game:     &entity.Game{ID: 1},
				genreIDs: []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s.GetByIDs
			genreRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).
				Return([]entity.Genre{{ID: 1}, {ID: 2}, {ID: 3}}, nil)

			// s.GameGenreRepository.FindByGameID
			gameGenreRepository.EXPECT().FindByGameID(gomock.Any(), gomock.Any()).
				Return([]entity.GameGenre{{GenreID: 3}}, nil)

			gameGenreRepository.EXPECT().DeleteMultiple(gomock.Any(), gomock.Any())
			gameGenreRepository.EXPECT().CreateMultiple(gomock.Any(), gomock.Any())

			err := s.UpdateGenresForGame(tt.args.ctx, tt.args.game, tt.args.genreIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateGenresForGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
