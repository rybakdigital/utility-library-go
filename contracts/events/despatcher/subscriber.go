package despatcher

type Subscriber interface {
	Process(m Message)
}
