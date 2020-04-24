package game_store

import (
	"encoding/json"

	"github.com/streadway/amqp"

	"github.com/google/uuid"
	"github.com/qilin/crm-api/internal/domain/publisher"
)

func (p Publisher) Publish(data publisher.PublishGameStoreData) error {
	genres := make([]publisher.GameStoreGenreMessage, len(data.Game.Revision.Genres))
	for _, item := range data.Game.Revision.Genres {
		genres = append(genres, publisher.GameStoreGenreMessage{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	tags := make([]publisher.GameStoreTagMessage, len(data.Game.Revision.Tags))
	for _, item := range data.Game.Revision.Tags {
		tags = append(tags, publisher.GameStoreTagMessage{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	features := make([]publisher.GameStoreFeatureMessage, len(data.Game.Revision.Features))
	for _, item := range data.Game.Revision.Features {
		features = append(features, publisher.GameStoreFeatureMessage{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	developers := make([]publisher.GameStoreDeveloperMessage, len(data.Game.Revision.Developers))
	for _, item := range data.Game.Revision.Developers {
		developers = append(developers, publisher.GameStoreDeveloperMessage{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	publishers := make([]publisher.GameStorePublisherMessage, len(data.Game.Revision.Publishers))
	for _, item := range data.Game.Revision.Publishers {
		publishers = append(publishers, publisher.GameStorePublisherMessage{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	message := &publisher.GameStoreMessage{
		ID:          data.Game.ID,
		Title:       data.Game.Title,
		Type:        data.Game.Type.String(),
		RevisionID:  data.Game.Revision.ID,
		Summary:     data.Game.Revision.Summary,
		Description: data.Game.Revision.Description,
		Slug:        data.Game.Revision.Slug,
		License:     data.Game.Revision.License,
		Platforms:   nil,
		ReleaseDate: data.Game.Revision.ReleaseDate,
		Tags:        tags,
		Features:    features,
		Developers:  developers,
		Publishers:  publishers,
		Genres:      genres,
		Rand:        uuid.New().String(),
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.Channel.Publish(ExchangeName, QueueName, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
}
