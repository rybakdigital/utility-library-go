package despatcher

type Adapter interface {
	Send(e Event, name string) (*Receipt, error)
}
