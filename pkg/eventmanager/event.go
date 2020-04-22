package eventmanager

import "go.uber.org/fx"

type Event struct {
	Topic string
}

type EventResult struct {
	fx.Out

	Event Event `group:"event"`
}
