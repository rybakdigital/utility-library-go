package despatcher

type Message interface {
	GetId() string
	GetData() []byte
	GetEventName() string
}
