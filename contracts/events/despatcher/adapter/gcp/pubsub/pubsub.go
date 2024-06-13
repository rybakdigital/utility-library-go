package pubsub

import (
	"context"
	"fmt"

	log "github.com/rybakdigital/utility-library-go/logging/logger"

	"cloud.google.com/go/pubsub"
	"github.com/rybakdigital/utility-library-go/contracts/events/despatcher"
)

const AdapterName = "GCPPubSub"

type PubSubAdapter struct {
	ProjectID string
	Prefix    string
	Logger    *log.Logger
}

type Receipt struct {
	Id string
}

type Event struct {
}

func (e *Event) GetData() []byte {
	return []byte("")
}

func (e *Event) GetName() string {
	return ""
}

func (e *Receipt) GetId() string {
	return ""
}

func New(projectID string, l *log.Logger) *PubSubAdapter {
	a := PubSubAdapter{
		ProjectID: projectID,
		Logger:    l,
	}

	l.Printf("Created new PubSub adapter")
	return &a
}

func (a *PubSubAdapter) GetName() string {
	return AdapterName
}

func (a *PubSubAdapter) Send(e despatcher.Event) (despatcher.Receipt, error) {
	a.Logger.Printf("%s: Received request to despatch the message for event %s", a.GetName(), e.GetEventId())

	// Set topic
	topicID := a.Prefix + e.GetName()
	a.Logger.Printf("Sending message to topic %s", topicID)

	// Create client
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, a.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)

	result := t.Publish(ctx, &pubsub.Message{
		Data: e.GetData(),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("pubsub: result.Get: %w", err)
	}

	a.Logger.Printf("Event %s published to topic %s: ID: %v", e.GetEventId(), topicID, id)

	return &Receipt{Id: id}, nil
}
