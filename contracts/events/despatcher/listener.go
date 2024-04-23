package despatcher

import (
	log "github.com/rybakdigital/utility-library-go/logging/logger"

	mapset "github.com/deckarep/golang-set/v2"
)

type ListenerAdapter interface {
	Receive(events chan Message)
}

type Listener struct {
	Logger      *log.Logger
	Subscribers mapset.Set[Subscriber]
	Adapters    mapset.Set[ListenerAdapter]
}

func NewListener(logger *log.Logger) *Listener {
	return &Listener{
		Subscribers: mapset.NewSet[Subscriber](),
		Adapters:    mapset.NewSet[ListenerAdapter](),
		Logger:      logger,
	}
}

func (l *Listener) Listen() {
	messages := make(chan Message)
	for _, adapter := range l.Adapters.ToSlice() {
		go adapter.Receive(messages)
	}

	for message := range messages {
		l.Logger.Printf("Received message %s: Message %s", message.GetId(), message.GetData())
	}
}
