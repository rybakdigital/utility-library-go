package despatcher

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	log "github.com/rybakdigital/utility-library-go/logging/logger"

	mapset "github.com/deckarep/golang-set/v2"
)

const DespatcherContextName = "event.despatcher"

type Despatcher struct {
	Logger   *log.Logger
	Adapters mapset.Set[Adapter]
	Listener *Listener
}

func New(testAdapters int, testMessages int) *Despatcher {
	// Create logger
	log := log.NewLogger("event-despatcher")
	list := NewListener(log)
	//list.MaxMessages = testMessages

	// Create new despatcher
	d := &Despatcher{
		Logger:   log,
		Adapters: mapset.NewSet[Adapter](),
		Listener: list,
		// Listener: NewListener(log),
	}

	// Log new despatcher
	d.Logger.Printf("Created new Event Despatcher")

	for i := 0; i < testAdapters; i++ {
		itner := rand.Intn(3)
		d.Listener.Adapters.Add(NewSimpleAdapter("Simple-"+strconv.Itoa(i), itner+1))
	}

	return d
}

func (d *Despatcher) Despatch(e Event, name string) error {
	// Get adapters and send the message
	for _, adapter := range d.Adapters.ToSlice() {
		receipt, err := adapter.Send(e, name)
		fmt.Println(receipt)
		fmt.Println(err)
	}

	return nil
}

func (d *Despatcher) Shutdown(ctx context.Context) {
	d.Logger.InfoF("Received request to shutdown")
	d.Listener.WaitForListenersToClose()
	d.Logger.InfoF("Shutdown complete")
}
