package despatcher

type Message interface {
	GetId() string        // Provider assigned ID
	GetData() []byte      // Message content in byte format, Payload in the byte format
	GetPayload() Payload  // Structured content
	GetEventName() string // Shortcut method to Payload.EventName
}

type Payload interface {
	GetEventId() string   // Unique ID of the event
	GetEventName() string // Name of the event
	GetContent() []byte   // Data sent by the event
}

type EventMessage struct {
	Id      string          `json:"id"`
	Payload *MessagePayload `json:"payload"`
}

type MessagePayload struct {
	EventId   string `json:"eventId"`
	EventName string `json:"eventName"`
	Content   []byte `json:"content"`
}

func (em *EventMessage) GetId() string {
	return em.Id
}

func (em *EventMessage) GetData() []byte {
	return em.Payload.Content
}

func (em *EventMessage) GetPayload() Payload {
	return em.Payload
}

func (em *EventMessage) GetEventName() string {
	return em.Payload.EventName
}

func (mp *MessagePayload) GetEventId() string {
	return mp.EventId
}

func (mp *MessagePayload) GetEventName() string {
	return mp.EventName
}

func (mp *MessagePayload) GetContent() []byte {
	return mp.Content
}
