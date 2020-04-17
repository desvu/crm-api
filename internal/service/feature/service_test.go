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

	type args struct {
		gameID              uint
		newFeatureIDs       []uint
		currentGameFeatures []entity.GameRevisionFeature
	}
	tests := []struct {
		name string
		args args
		want []entity.GameRevisionFeature
	}{
		{
			name: "getting a list of feature IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:              1,
				newFeatureIDs:       []uint{1, 2, 3, 4},
				currentGameFeatures: []entity.GameRevisionFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameRevisionFeature{{FeatureID: 1, GameRevisionID: 1}, {FeatureID: 4, GameRevisionID: 1}},
		},
		{
			name: "getting a list of feature IDs with a fully included subset of IDs associated with the device",
			args: args{
				gameID:              1,
				newFeatureIDs:       []uint{2, 3},
				currentGameFeatures: []entity.GameRevisionFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameRevisionFeature{},
		},
		{
			name: "getting a list of feature IDs with or without an incoming subset of IDs associated with the device",
			args: args{
				gameID:              1,
				newFeatureIDs:       []uint{5, 6},
				currentGameFeatures: []entity.GameRevisionFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameRevisionFeature{{FeatureID: 5, GameRevisionID: 1}, {FeatureID: 6, GameRevisionID: 1}},
		},
		{
			name: "getting a list of feature IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:              1,
				newFeatureIDs:       []uint{1, 2, 2, 3, 4},
				currentGameFeatures: []entity.GameRevisionFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameRevisionFeature{{FeatureID: 1, GameRevisionID: 1}, {FeatureID: 4, GameRevisionID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameFeaturesForInsert(tt.args.gameID, tt.args.newFeatureIDs, tt.args.currentGameFeatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameFeaturesForInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_getGameFeaturesForDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		newFeatureIDs       []uint
		currentGameFeatures []entity.GameRevisionFeature
	}
	tests := []struct {
		name string
		args args
		want []entity.GameRevisionFeature
	}{
		{
			name: "getting a list of feature IDs with a partially included subset of IDs associated with the device",
			args: args{
				newFeatureIDs:       []uint{1, 2, 3, 4},
				currentGameFeatures: []entity.GameRevisionFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameRevisionFeature{},
		},
		{
			name: "getting a list of feature IDs with a fully included subset of IDs associated with the device",
			args: args{
				newFeatureIDs:       []uint{2, 3},
				currentGameFeatures: []entity.GameRevisionFeature{{FeatureID: 2}, {FeatureID: 3}},
			},
			want: []entity.GameRevisionFeature{},
		},
		{
			name: "getting a list of feature IDs with a partially included subset of IDs associated with the device",
			args: args{
				newFeatureIDs:       []uint{1, 4},
				currentGameFeatures: []entity.GameRevisionFeature{{ID: 1, FeatureID: 2, GameRevisionID: 1}, {ID: 1, FeatureID: 3, GameRevisionID: 1}},
			},
			want: []entity.GameRevisionFeature{{ID: 1, FeatureID: 2, GameRevisionID: 1}, {ID: 1, FeatureID: 3, GameRevisionID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameFeaturesForDelete(tt.args.newFeatureIDs, tt.args.currentGameFeatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameFeaturesForDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateFeaturesForGame(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	featureRepository := mocks.NewMockFeatureRepository(ctrl)
	gameRevisionFeatureRepository := mocks.NewMockGameRevisionFeatureRepository(ctrl)

	s := New(ServiceParams{
		FeatureRepository:             featureRepository,
		GameRevisionFeatureRepository: gameRevisionFeatureRepository,
	})

	type args struct {
		ctx          context.Context
		gameRevision *entity.GameRevision
		featureIDs   []uint
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
				ctx:          context.Background(),
				gameRevision: &entity.GameRevision{ID: 1},
				featureIDs:   []uint{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "getting a non-existent feature ID",
			args: args{
				ctx:          context.Background(),
				gameRevision: &entity.GameRevision{ID: 1},
				featureIDs:   []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s.GetByIDs
			featureRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).
				Return([]entity.Feature{{ID: 1}, {ID: 2}, {ID: 3}}, nil)

			// s.GameRevisionFeatureRepository.FindByGameRevisionID
			gameRevisionFeatureRepository.EXPECT().FindByGameRevisionID(gomock.Any(), gomock.Any()).
				Return([]entity.GameRevisionFeature{{FeatureID: 3}}, nil)

			gameRevisionFeatureRepository.EXPECT().DeleteMultiple(gomock.Any(), gomock.Any())
			gameRevisionFeatureRepository.EXPECT().CreateMultiple(gomock.Any(), gomock.Any())

			err := s.UpdateFeaturesForGameRevision(tt.args.ctx, tt.args.gameRevision, tt.args.featureIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateFeaturesForGameRevision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
