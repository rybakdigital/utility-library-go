package despatcher

type Adapter interface {
	GetName() string
	Send(e Event, name string) (Receipt, error)
}
