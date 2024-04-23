package despatcher

import (
	"fmt"
	"strconv"
	"time"
)

type SimpleAdapter struct {
	Name     string
	Interval int
}

type SimpleMessage struct {
	Id   string
	Data string
}

func (m *SimpleMessage) GetId() string {
	return m.Id
}

func (m *SimpleMessage) GetData() []byte {
	return []byte(m.Data)
}

func NewSimpleAdapter(name string, interval int) *SimpleAdapter {
	return &SimpleAdapter{Name: name, Interval: interval}
}

func (a *SimpleAdapter) Receive(events chan Message) {
	i := 0
	for {
		time.Sleep(time.Second * time.Duration(a.Interval))
		msg := &SimpleMessage{
			Id:   strconv.Itoa(i),
			Data: fmt.Sprintf("Adapter %s: Message received", a.Name),
		}
		events <- msg
		i++
	}
}
