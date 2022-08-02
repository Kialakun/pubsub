package pubsub

import "sync"

type Event struct {
	Data  interface{} // store anything
	Topic string      // Topic to publish to
}

type EventChannel chan Event

type EventChannels []EventChannel

type EventBus struct {
	subscribers map[string]EventChannels // map of Topic to EventChannels
	mutx        sync.Mutex               // Mutex to safely read/write
}
