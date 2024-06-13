package despatcher

type Event interface {
	GetData() []byte
	GetEventId() string
	GetName() string
}
