package despatcher

import (
	"context"
	"sync"
	"time"

	log "github.com/rybakdigital/utility-library-go/logging/logger"

	mapset "github.com/deckarep/golang-set/v2"
)

// ListenerAdapter is a receiver. It receives messages from the source (PuSub, Queue, Redis, etc)
// and then forwards these messages to the Listener
type ListenerAdapter interface {
	Receive(events chan Message, stopCh chan bool, feedbackCh chan bool)
}

// Listener manages inflow of the messages. It registers ListenerAdapter (Receivers) and receives messages
// from these adapters. It then decides which Subscribers should receive these messages. It also manages Adapters
// by telling them to stop process any more messages.
type Listener struct {
	Logger          *log.Logger
	Subscribers     mapset.Set[Subscriber]
	Adapters        mapset.Set[ListenerAdapter]
	ReceiveMessages bool
	Listens         bool
}

func NewListener(logger *log.Logger) *Listener {
	return &Listener{
		Subscribers: mapset.NewSet[Subscriber](),
		Adapters:    mapset.NewSet[ListenerAdapter](),
		Logger:      logger,
	}
}

func (l *Listener) Listen(ctx context.Context) {
	l.Listens = true // Set Listener to listens mode

	// Create comm channels for Adapters
	messages := make(chan Message)
	stopCh := make(chan bool)
	feedbackCh := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	// Monitor number of active Adapters (Receivers)
	receivers := 0

	// Tell Adapters to start receiving messages
	for _, adapter := range l.Adapters.ToSlice() {
		go adapter.Receive(messages, stopCh, feedbackCh)
		receivers++
	}

	// We can now start processing incomming messages
	go func() {
		i := 0
		defer wg.Done()
		for receivers > 0 {
			select {
			case message := <-messages:
				l.Logger.Printf("Received message %d: Message ID %s: %s", i, message.GetId(), message.GetData())
				i++
			case <-feedbackCh:
				// Receiver has been evicted from the pool
				l.Logger.Printf("Receiver has stopped receiving messages")
				receivers -= 1
				l.Logger.Printf("Still got %d active receivers", receivers)

				// Keep message receiver open until all receivers are shut down
				// This is to avoid unprocessed messages being lost
				if receivers == 0 {
					close(messages)
					l.Logger.Printf("All receivers stopped receiving messages. Closing channel")
					return
				}
			case <-ctx.Done():
				// We have received signal to stop receiving further messages
				// Let's inform Receivers by sending signal on the stopCh channel
				if !l.ReceiveMessages {
					close(stopCh)
				}
				l.ReceiveMessages = true
			}
		}
	}()

	// Wait for all messages to be received
	wg.Wait()

	// Set Listener to inactive mode
	l.Listens = false
	l.Logger.Printf("All messages have been processed. There are %d active receivers. Closing listener", receivers)
}

// WaitForListenersToClose for all Adapters to stop sending messages
// This is blocking function
func (l *Listener) WaitForListenersToClose() {
	for {
		if !l.Listens {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
