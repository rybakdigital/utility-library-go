package despatcher

type Payload interface {
	GetEventId() string   // Unique ID of the event
	GetEventName() string // Name of the event
	GetContent() []byte   // Data sent by the event
}
