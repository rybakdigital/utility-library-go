package despatcher

type Adapter interface {
	GetName() string
	Send(m Message) (Receipt, error)
}
