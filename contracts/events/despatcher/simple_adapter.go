package despatcher

import (
	"fmt"
	"log"
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

func (a *SimpleAdapter) Receive(events chan Message, stopCh chan bool, feedbackCh chan bool) {
	i := 0
	for {
		select {
		case <-stopCh:
			log.Printf("Adapter [%s]: Received request to stop the messages", a.Name)
			feedbackCh <- true
			return
		default:
		}
		time.Sleep(time.Second * time.Duration(a.Interval))
		msg := &SimpleMessage{
			Id:   strconv.Itoa(i),
			Data: fmt.Sprintf("Adapter [%s]: Message received", a.Name),
		}

		log.Printf("Adapter [%s]: Sending message: %d", a.Name, i)
		events <- msg
		i++
	}
}
