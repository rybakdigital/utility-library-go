package despatcher

import mapset "github.com/deckarep/golang-set/v2"

type Subscriber interface {
	GetName() string
	Process(m Message)
	GetSubscribedEvents() mapset.Set[string]
}
