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
	Subscriptions   map[string][]Subscriber
	ReceiveMessages bool
	Listens         bool
}

func NewListener(logger *log.Logger) *Listener {
	return &Listener{
		Subscribers:   mapset.NewSet[Subscriber](),
		Adapters:      mapset.NewSet[ListenerAdapter](),
		Subscriptions: map[string][]Subscriber{},
		Logger:        logger,
	}
}

func (l *Listener) Listen(ctx context.Context) {
	l.Listens = true // Set Listener to listens mode

	// Create comm channels for Adapters
	messages := make(chan Message)
	stopCh := make(chan bool)
	feedbackCh := make(chan bool)

	var wg sync.WaitGroup
	msgCounter := &sync.WaitGroup{}

	// Monitor number of active Adapters (Receivers)
	receivers := 0

	// Monitor number of messages processed
	messagesToBeProcessedCount := 0

	// Tell Adapters to start receiving messages
	for _, adapter := range l.Adapters.ToSlice() {
		go adapter.Receive(messages, stopCh, feedbackCh)
		receivers++
	}

	// We can now start processing incoming messages
	wg.Add(1)
	go func() {
		defer wg.Done()
		for receivers > 0 {
			select {
			case message := <-messages:
				l.Logger.Printf("Received message for event type %s : Message ID %s: %s", message.GetEventName(), message.GetId(), message.GetData())

				// Find subscriptions for this event
				l.Logger.Printf("Message ID %s. Checking for subscribers to event %s", message.GetId(), message.GetEventName())
				if subscriptions, ok := l.Subscriptions[message.GetEventName()]; ok {
					for _, subscriber := range subscriptions {
						// Found subscriber for this event
						l.Logger.Printf("Message ID %s. Subscriber %s has active subscription to event %s", message.GetId(), subscriber.GetName(), message.GetEventName())

						// Increase messages counter to track processed messages
						msgCounter.Add(1)

						// Forward message to Subscriber
						l.Logger.Printf("Forwarding Message ID %s for event %s to subscriber %s", message.GetId(), message.GetEventName(), subscriber.GetName())
						go func(subscriber Subscriber) {
							defer msgCounter.Done()

							// Forward message
							subscriber.Process(message)
						}(subscriber)
						messagesToBeProcessedCount++
					}
				}
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
	msgCounter.Wait()

	// Set Listener to inactive mode
	l.Listens = false
	l.Logger.Printf("All messages have been processed. There are %d active receivers. Closing listener", receivers)
	l.Logger.Printf("Total messages send for processing: %d", messagesToBeProcessedCount)
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

func (l *Listener) AddSubscription(eventName string, subscriber Subscriber) {
	l.Subscriptions[eventName] = append(l.Subscriptions[eventName], subscriber)
}
