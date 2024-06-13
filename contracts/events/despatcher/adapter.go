package despatcher

type Adapter interface {
	GetName() string
	Send(e Event) (Receipt, error)
}
