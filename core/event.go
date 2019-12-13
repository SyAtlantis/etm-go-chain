package core

import "sync"

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel 是一个能接收 DataEvent 的 channel
type DataChannel chan DataEvent

// DataChannelSlice 是一个包含 DataChannels 数据的切片
type DataChannelSlice []DataChannel

type Event struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

var event *Event

func NewEvent() *Event {
	if event == nil {
		event = &Event{
			subscribers: make(map[string]DataChannelSlice),
		}
	}
	return event
}

func (e *Event) Subscribe(topic string, ch DataChannel) {
	e.rm.Lock()

	if prev, found := e.subscribers[topic]; found {
		e.subscribers[topic] = append(prev, ch)
	} else {
		e.subscribers[topic] = append([]DataChannel{}, ch)
	}

	e.rm.Unlock()
}

func (e *Event) Publish(topic string, data interface{}) {
	e.rm.RLock()

	if chs, found := e.subscribers[topic]; found {
		channels := append(DataChannelSlice{}, chs...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}

	e.rm.RUnlock()
}
