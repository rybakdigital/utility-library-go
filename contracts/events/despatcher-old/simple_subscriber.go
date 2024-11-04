package despatcher

import (
	"log"
	"math/rand"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

type SimpleSubscriber struct {
	Name   string
	Events mapset.Set[string]
}

func (s *SimpleSubscriber) Process(m Message) {
	itner := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(itner))
	log.Printf("Subscriber %s - Processing message %s", s.Name, m.GetId())
}

func (s *SimpleSubscriber) GetName() string {
	return s.Name
}

func (s *SimpleSubscriber) GetSubscribedEvents() mapset.Set[string] {
	return s.Events
}

func NewSimpleSubscriber(name string, events []string) *SimpleSubscriber {
	return &SimpleSubscriber{Name: name, Events: mapset.NewSet[string](events...)}
}
