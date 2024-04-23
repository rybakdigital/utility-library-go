package despatcher

import (
	"sync"

	log "github.com/rybakdigital/utility-library-go/logging/logger"

	mapset "github.com/deckarep/golang-set/v2"
)

type ListenerAdapter interface {
	Receive(events chan Message, stopCh chan bool, feedbackCh chan bool)
}

type Listener struct {
	Logger      *log.Logger
	Subscribers mapset.Set[Subscriber]
	Adapters    mapset.Set[ListenerAdapter]
	MaxMessages int
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
	stopCh := make(chan bool)
	feedbackCh := make(chan bool)
	senders := 0
	var wg sync.WaitGroup
	wg.Add(1)
	for _, adapter := range l.Adapters.ToSlice() {
		go adapter.Receive(messages, stopCh, feedbackCh)
		senders++
	}

	go func() {
		i := 0
		defer wg.Done()
		for {
			select {
			case message := <-messages:
				l.Logger.Printf("Received message %d: Message ID %s: %s", i, message.GetId(), message.GetData())
				i++

				if i == l.MaxMessages {
					close(stopCh)
				}
			case <-feedbackCh:
				l.Logger.Printf("Sender stopped sending messages")
				senders -= 1
				l.Logger.Printf("Still go %d active senders", senders)

				if senders == 0 {
					close(messages)
					l.Logger.Printf("All senders stopped sending messages. Closing channel")
					return
				}
			}
		}
	}()

	wg.Wait()
	l.Logger.Printf("All messages have been processed. Closing listener")
}
