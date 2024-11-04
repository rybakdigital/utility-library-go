package despatcher

type Message interface {
	GetId() string        // Provider assigned ID
	GetPayload() Payload  // Structured content
	GetEventName() string // Shortcut method to Payload.EventName
}
