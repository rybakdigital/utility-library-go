package despatcher

import (
	"encoding/json"
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
	Id      string
	Payload *SimplePayload
}

type SimplePayload struct {
	Content   []byte `json:"content"`
	EventId   string `json:"eventId"`
	EventName string `json:"eventName"`
}

func (m *SimpleMessage) GetId() string {
	return m.Id
}

func (m *SimpleMessage) GetData() []byte {
	data, _ := json.Marshal(&m.Payload)
	fmt.Println(data)
	return data
}

func (m *SimpleMessage) GetEventName() string {
	return m.Payload.EventName
}

func (m *SimpleMessage) GetPayload() Payload {
	return m.Payload
}

func (p *SimplePayload) GetEventName() string {
	return p.EventName
}

func (p *SimplePayload) GetEventId() string {
	return p.EventId
}

func (p *SimplePayload) GetContent() []byte {
	return p.Content
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
			Id: strconv.Itoa(i),
			Payload: &SimplePayload{
				EventName: "user.created",
				Content:   []byte(fmt.Sprintf("Adapter [%s]: Message received", a.Name)),
			},
		}

		log.Printf("Adapter [%s]: Sending message: %d", a.Name, i)
		events <- msg
		i++
	}
}
