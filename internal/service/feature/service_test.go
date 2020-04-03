package feature

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/mocks"
)

func TestService_getGameFeaturesForInsert(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	gameService := mocks.NewMockGameService(ctrl)
	featureRepository := mocks.NewMockFeatureRepository(ctrl)
	gameFeatureRepository := mocks.NewMockGameFeatureRepository(ctrl)

	s := New(ServiceParams{
		GameService:           gameService,
		FeatureRepository:     featureRepository,
		GameFeatureRepository: gameFeatureRepository,
	})

	type args struct {
		gameID              uint
		newFeatureIDs       []uint
		currentGameFeatures []entity.GameFeature
	}
	tests := []struct {
		name string
		args args
		want []entity.GameFeature
	}{
		{
			name: "getting a list of feature IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:              1,
				newFeatureIDs:       []uint{1, 2, 3, 4},
				currentGameFeatures: []entity.GameFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameFeature{{FeatureID: 1, GameID: 1}, {FeatureID: 4, GameID: 1}},
		},
		{
			name: "getting a list of feature IDs with a fully included subset of IDs associated with the device",
			args: args{
				gameID:              1,
				newFeatureIDs:       []uint{2, 3},
				currentGameFeatures: []entity.GameFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameFeature{},
		},
		{
			name: "getting a list of feature IDs with or without an incoming subset of IDs associated with the device",
			args: args{
				gameID:              1,
				newFeatureIDs:       []uint{5, 6},
				currentGameFeatures: []entity.GameFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameFeature{{FeatureID: 5, GameID: 1}, {FeatureID: 6, GameID: 1}},
		},
		{
			name: "getting a list of feature IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:              1,
				newFeatureIDs:       []uint{1, 2, 2, 3, 4},
				currentGameFeatures: []entity.GameFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameFeature{{FeatureID: 1, GameID: 1}, {FeatureID: 4, GameID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.getGameFeaturesForInsert(tt.args.gameID, tt.args.newFeatureIDs, tt.args.currentGameFeatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameFeaturesForInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_getGameFeaturesForDelete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	gameService := mocks.NewMockGameService(ctrl)
	featureRepository := mocks.NewMockFeatureRepository(ctrl)
	gameFeatureRepository := mocks.NewMockGameFeatureRepository(ctrl)

	s := New(ServiceParams{
		GameService:           gameService,
		FeatureRepository:     featureRepository,
		GameFeatureRepository: gameFeatureRepository,
	})

	type args struct {
		newFeatureIDs       []uint
		currentGameFeatures []entity.GameFeature
	}
	tests := []struct {
		name string
		args args
		want []entity.GameFeature
	}{
		{
			name: "getting a list of feature IDs with a partially included subset of IDs associated with the device",
			args: args{
				newFeatureIDs:       []uint{1, 2, 3, 4},
				currentGameFeatures: []entity.GameFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameFeature{},
		},
		{
			name: "getting a list of feature IDs with a fully included subset of IDs associated with the device",
			args: args{
				newFeatureIDs:       []uint{2, 3},
				currentGameFeatures: []entity.GameFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameFeature{},
		},
		{
			name: "getting a list of feature IDs with a partially included subset of IDs associated with the device",
			args: args{
				newFeatureIDs:       []uint{1, 4},
				currentGameFeatures: []entity.GameFeature{{ID: 1, FeatureID: 2, GameID: 1}, {ID: 1, FeatureID: 3, GameID: 1}},
			},
			want: []entity.GameFeature{{ID: 1, FeatureID: 2, GameID: 1}, {ID: 1, FeatureID: 3, GameID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.getGameFeaturesForDelete(tt.args.newFeatureIDs, tt.args.currentGameFeatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameFeaturesForDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateFeaturesForGame(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	gameService := mocks.NewMockGameService(ctrl)
	featureRepository := mocks.NewMockFeatureRepository(ctrl)
	gameFeatureRepository := mocks.NewMockGameFeatureRepository(ctrl)

	s := New(ServiceParams{
		GameService:           gameService,
		FeatureRepository:     featureRepository,
		GameFeatureRepository: gameFeatureRepository,
	})

	type args struct {
		ctx        context.Context
		game       *entity.Game
		featureIDs []uint
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Feature
		wantErr bool
	}{
		{
			name: "getting a non-existent feature ID",
			args: args{
				ctx:        context.Background(),
				game:       &entity.Game{ID: 1},
				featureIDs: []uint{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "getting a non-existent feature ID",
			args: args{
				ctx:        context.Background(),
				game:       &entity.Game{ID: 1},
				featureIDs: []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s.GetByIDs
			featureRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).
				Return([]entity.Feature{{ID: 1}, {ID: 2}, {ID: 3}}, nil)

			// s.GameFeatureRepository.FindByGameID
			gameFeatureRepository.EXPECT().FindByGameID(gomock.Any(), gomock.Any()).
				Return([]entity.GameFeature{{FeatureID: 3}}, nil)

			gameFeatureRepository.EXPECT().DeleteMultiple(gomock.Any(), gomock.Any())
			gameFeatureRepository.EXPECT().CreateMultiple(gomock.Any(), gomock.Any())

			err := s.UpdateFeaturesForGame(tt.args.ctx, tt.args.game, tt.args.featureIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateFeaturesForGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
