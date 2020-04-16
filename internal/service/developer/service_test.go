package developer

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/mocks"
)

func TestService_getGameDevelopersForInsert(t *testing.T) {
	t.Parallel()

	type args struct {
		gameID                uint
		newDeveloperIDs       []uint
		currentGameDevelopers []entity.GameRevisionDeveloper
	}
	tests := []struct {
		name string
		args args
		want []entity.GameRevisionDeveloper
	}{
		{
			name: "getting a list of developer IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:                1,
				newDeveloperIDs:       []uint{1, 2, 3, 4},
				currentGameDevelopers: []entity.GameRevisionDeveloper{{DeveloperID: 2}, {DeveloperID: 3}},
			},
			want: []entity.GameRevisionDeveloper{{DeveloperID: 1, GameRevisionID: 1}, {DeveloperID: 4, GameRevisionID: 1}},
		},
		{
			name: "getting a list of developer IDs with a fully included subset of IDs associated with the device",
			args: args{
				gameID:                1,
				newDeveloperIDs:       []uint{2, 3},
				currentGameDevelopers: []entity.GameRevisionDeveloper{{DeveloperID: 2}, {DeveloperID: 3}},
			},
			want: []entity.GameRevisionDeveloper{},
		},
		{
			name: "getting a list of developer IDs with or without an incoming subset of IDs associated with the device",
			args: args{
				gameID:                1,
				newDeveloperIDs:       []uint{5, 6},
				currentGameDevelopers: []entity.GameRevisionDeveloper{{DeveloperID: 2}, {DeveloperID: 3}},
			},
			want: []entity.GameRevisionDeveloper{{DeveloperID: 5, GameRevisionID: 1}, {DeveloperID: 6, GameRevisionID: 1}},
		},
		{
			name: "getting a list of developer IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:                1,
				newDeveloperIDs:       []uint{1, 2, 2, 3, 4},
				currentGameDevelopers: []entity.GameRevisionDeveloper{{DeveloperID: 2}, {DeveloperID: 3}},
			},
			want: []entity.GameRevisionDeveloper{{DeveloperID: 1, GameRevisionID: 1}, {DeveloperID: 4, GameRevisionID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameDevelopersForInsert(tt.args.gameID, tt.args.newDeveloperIDs, tt.args.currentGameDevelopers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameDevelopersForInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_getGameDevelopersForDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		newDeveloperIDs       []uint
		currentGameDevelopers []entity.GameRevisionDeveloper
	}
	tests := []struct {
		name string
		args args
		want []entity.GameRevisionDeveloper
	}{
		{
			name: "getting a list of developer IDs with a partially included subset of IDs associated with the device",
			args: args{
				newDeveloperIDs:       []uint{1, 2, 3, 4},
				currentGameDevelopers: []entity.GameRevisionDeveloper{{DeveloperID: 2}, {DeveloperID: 3}},
			},
			want: []entity.GameRevisionDeveloper{},
		},
		{
			name: "getting a list of developer IDs with a fully included subset of IDs associated with the device",
			args: args{
				newDeveloperIDs:       []uint{2, 3},
				currentGameDevelopers: []entity.GameRevisionDeveloper{{DeveloperID: 2}, {DeveloperID: 3}},
			},
			want: []entity.GameRevisionDeveloper{},
		},
		{
			name: "getting a list of developer IDs with a partially included subset of IDs associated with the device",
			args: args{
				newDeveloperIDs:       []uint{1, 4},
				currentGameDevelopers: []entity.GameRevisionDeveloper{{ID: 1, DeveloperID: 2, GameRevisionID: 1}, {ID: 1, DeveloperID: 3, GameRevisionID: 1}},
			},
			want: []entity.GameRevisionDeveloper{{ID: 1, DeveloperID: 2, GameRevisionID: 1}, {ID: 1, DeveloperID: 3, GameRevisionID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameDevelopersForDelete(tt.args.newDeveloperIDs, tt.args.currentGameDevelopers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGameDevelopersForDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateDevelopersForGame(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	developerRepository := mocks.NewMockDeveloperRepository(ctrl)
	gameRevisionDeveloperRepository := mocks.NewMockGameRevisionDeveloperRepository(ctrl)

	s := New(ServiceParams{
		DeveloperRepository:             developerRepository,
		GameRevisionDeveloperRepository: gameRevisionDeveloperRepository,
	})

	type args struct {
		ctx          context.Context
		game         *entity.GameRevision
		developerIDs []uint
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Developer
		wantErr bool
	}{
		{
			name: "getting a non-existent developer ID",
			args: args{
				ctx:          context.Background(),
				game:         &entity.GameRevision{ID: 1},
				developerIDs: []uint{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "getting a non-existent developer ID",
			args: args{
				ctx:          context.Background(),
				game:         &entity.GameRevision{ID: 1},
				developerIDs: []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s.GetByIDs
			developerRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).
				Return([]entity.Developer{{ID: 1}, {ID: 2}, {ID: 3}}, nil)

			// s.GameRevisionDeveloperRepository.FindByGameRevisionID
			gameRevisionDeveloperRepository.EXPECT().FindByGameRevisionID(gomock.Any(), gomock.Any()).
				Return([]entity.GameRevisionDeveloper{{DeveloperID: 3}}, nil)

			gameRevisionDeveloperRepository.EXPECT().DeleteMultiple(gomock.Any(), gomock.Any())
			gameRevisionDeveloperRepository.EXPECT().CreateMultiple(gomock.Any(), gomock.Any())

			err := s.UpdateDevelopersForGameRevision(tt.args.ctx, tt.args.game, tt.args.developerIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateDevelopersForGameRevision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
