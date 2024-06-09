package despatcher

type Event interface {
	GetData() []byte
	GetName() string
}
