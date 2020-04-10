package publisher

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/mocks"
)

func TestService_getGamePublishersForInsert(t *testing.T) {
	t.Parallel()

	type args struct {
		gameID                uint
		newPublisherIDs       []uint
		currentGamePublishers []entity.GamePublisher
	}
	tests := []struct {
		name string
		args args
		want []entity.GamePublisher
	}{
		{
			name: "getting a list of publisher IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:                1,
				newPublisherIDs:       []uint{1, 2, 3, 4},
				currentGamePublishers: []entity.GamePublisher{{PublisherID: 2}, {PublisherID: 3}},
			},
			want: []entity.GamePublisher{{PublisherID: 1, GameID: 1}, {PublisherID: 4, GameID: 1}},
		},
		{
			name: "getting a list of publisher IDs with a fully included subset of IDs associated with the device",
			args: args{
				gameID:                1,
				newPublisherIDs:       []uint{2, 3},
				currentGamePublishers: []entity.GamePublisher{{PublisherID: 2}, {PublisherID: 3}},
			},
			want: []entity.GamePublisher{},
		},
		{
			name: "getting a list of publisher IDs with or without an incoming subset of IDs associated with the device",
			args: args{
				gameID:                1,
				newPublisherIDs:       []uint{5, 6},
				currentGamePublishers: []entity.GamePublisher{{PublisherID: 2}, {PublisherID: 3}},
			},
			want: []entity.GamePublisher{{PublisherID: 5, GameID: 1}, {PublisherID: 6, GameID: 1}},
		},
		{
			name: "getting a list of publisher IDs with a partially included subset of IDs associated with the device",
			args: args{
				gameID:                1,
				newPublisherIDs:       []uint{1, 2, 2, 3, 4},
				currentGamePublishers: []entity.GamePublisher{{PublisherID: 2}, {PublisherID: 3}},
			},
			want: []entity.GamePublisher{{PublisherID: 1, GameID: 1}, {PublisherID: 4, GameID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGamePublishersForInsert(tt.args.gameID, tt.args.newPublisherIDs, tt.args.currentGamePublishers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGamePublishersForInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_getGamePublishersForDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		newPublisherIDs       []uint
		currentGamePublishers []entity.GamePublisher
	}
	tests := []struct {
		name string
		args args
		want []entity.GamePublisher
	}{
		{
			name: "getting a list of publisher IDs with a partially included subset of IDs associated with the device",
			args: args{
				newPublisherIDs:       []uint{1, 2, 3, 4},
				currentGamePublishers: []entity.GamePublisher{{PublisherID: 2}, {PublisherID: 3}},
			},
			want: []entity.GamePublisher{},
		},
		{
			name: "getting a list of publisher IDs with a fully included subset of IDs associated with the device",
			args: args{
				newPublisherIDs:       []uint{2, 3},
				currentGamePublishers: []entity.GamePublisher{{PublisherID: 2}, {PublisherID: 3}},
			},
			want: []entity.GamePublisher{},
		},
		{
			name: "getting a list of publisher IDs with a partially included subset of IDs associated with the device",
			args: args{
				newPublisherIDs:       []uint{1, 4},
				currentGamePublishers: []entity.GamePublisher{{ID: 1, PublisherID: 2, GameID: 1}, {ID: 1, PublisherID: 3, GameID: 1}},
			},
			want: []entity.GamePublisher{{ID: 1, PublisherID: 2, GameID: 1}, {ID: 1, PublisherID: 3, GameID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGamePublishersForDelete(tt.args.newPublisherIDs, tt.args.currentGamePublishers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGamePublishersForDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdatePublishersForGame(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	publisherRepository := mocks.NewMockPublisherRepository(ctrl)
	gamePublisherRepository := mocks.NewMockGamePublisherRepository(ctrl)

	s := New(ServiceParams{
		PublisherRepository:     publisherRepository,
		GamePublisherRepository: gamePublisherRepository,
	})

	type args struct {
		ctx          context.Context
		game         *entity.Game
		publisherIDs []uint
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Publisher
		wantErr bool
	}{
		{
			name: "getting a non-existent publisher ID",
			args: args{
				ctx:          context.Background(),
				game:         &entity.Game{ID: 1},
				publisherIDs: []uint{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "getting a non-existent publisher ID",
			args: args{
				ctx:          context.Background(),
				game:         &entity.Game{ID: 1},
				publisherIDs: []uint{1, 2, 3, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s.GetByIDs
			publisherRepository.EXPECT().FindByIDs(gomock.Any(), gomock.Any()).
				Return([]entity.Publisher{{ID: 1}, {ID: 2}, {ID: 3}}, nil)

			// s.GamePublisherRepository.FindByGameID
			gamePublisherRepository.EXPECT().FindByGameID(gomock.Any(), gomock.Any()).
				Return([]entity.GamePublisher{{PublisherID: 3}}, nil)

			gamePublisherRepository.EXPECT().DeleteMultiple(gomock.Any(), gomock.Any())
			gamePublisherRepository.EXPECT().CreateMultiple(gomock.Any(), gomock.Any())

			err := s.UpdatePublishersForGame(tt.args.ctx, tt.args.game, tt.args.publisherIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePublishersForGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
