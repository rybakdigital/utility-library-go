package despatcher

import (
	"context"
	"errors"
	"fmt"

	log "github.com/rybakdigital/utility-library-go/logging/logger"

	mapset "github.com/deckarep/golang-set/v2"
)

const DespatcherContextName = "event.despatcher"

type Despatcher struct {
	Logger   *log.Logger
	Adapters mapset.Set[Adapter]
	Listener *Listener
}

func New(testAdapters int, testMessages int, testSubscribers int) *Despatcher {
	// Create logger
	log := log.NewLogger("event-despatcher")

	// Create new despatcher
	d := &Despatcher{
		Logger:   log,
		Adapters: mapset.NewSet[Adapter](),
		Listener: NewListener(log),
	}

	// Log new despatcher
	d.Logger.Printf("Created new Event Despatcher")

	// for i := 0; i < testAdapters; i++ {
	// 	itner := rand.Intn(3)
	// 	d.Listener.Adapters.Add(NewSimpleAdapter("Simple-"+strconv.Itoa(i), itner+1))
	// }

	// for i := 0; i < testSubscribers; i++ {
	// 	events := []string{"user.deleted"}
	// 	if i > 0 {
	// 		events = append(events, "user.created")
	// 	}

	// 	subscriber := NewSimpleSubscriber("Simple-"+strconv.Itoa(i), events)
	// 	d.Listener.AddSubscriber(subscriber)
	// }

	return d
}

func (d *Despatcher) Despatch(e Event) ([]Receipt, error) {
	d.Logger.Printf("Despatcher received request to despatch event %s with ID: %s", e.GetName(), e.GetEventId())

	// Get adapters and send the message
	if d.Adapters.IsEmpty() {
		msg := fmt.Sprintf("Tried to despatch event %s but no adapters were configured", e.GetEventId())
		d.Logger.Printf(msg)
		return nil, errors.New(msg)
	}

	var receipts []Receipt
	for _, adapter := range d.Adapters.ToSlice() {
		receipt, err := adapter.Send(e)

		if err != nil {
			msg := fmt.Sprintf("Error when sending message using adapter %s: %v", adapter.GetName(), err)
			d.Logger.Printf(msg)
			return nil, err
		}

		d.Logger.Printf("Adapter %s has sent the message. Receipt ID: %s", adapter.GetName(), receipt.GetId())
	}

	return receipts, nil
}

func (d *Despatcher) Shutdown(ctx context.Context) {
	d.Logger.InfoF("Received request to shutdown")
	d.Listener.WaitForListenersToClose()
	d.Logger.InfoF("Shutdown complete")
}
