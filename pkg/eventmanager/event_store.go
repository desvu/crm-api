package eventmanager

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/micro/go-micro/v2"
	"go.uber.org/fx"
)

var ErrEventNotFound = errors.New("event not found")

type EventStore struct {
	service micro.Service
	events  map[string]micro.Event
	mx      sync.RWMutex
}

type Params struct {
	fx.In

	Service micro.Service
	Events  []Event `group:"event"`
}

func New(params Params) *EventStore {
	es := &EventStore{
		service: params.Service,
		events:  make(map[string]micro.Event),
	}

	for _, event := range params.Events {
		es.Add(event)
	}

	return es
}

func (es *EventStore) Add(event Event) {
	es.mx.Lock()

	log.Println(event.Topic, es.service.Client())
	es.events[event.Topic] = micro.NewEvent(event.Topic, es.service.Client())

	es.mx.Unlock()
}

func (es *EventStore) Publish(ctx context.Context, topic string, data interface{}) error {
	es.mx.RLock()

	event, ok := es.events[topic]
	if !ok {
		return ErrEventNotFound
	}

	es.mx.RUnlock()

	return event.Publish(ctx, data)
}
