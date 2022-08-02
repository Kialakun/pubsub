package pubsub

// publishes data to a topic by sending it to all registered subscriber channels
func (e *EventBus) Publish(topic string, data interface{}) {
	// @todo: implement mutext to safely read/write from EventBus
	if subscribers, found := e.subscribers[topic]; found {
		for _, subscriber := range subscribers {
			subscriber <- Event{Data: data, Topic: topic}
		}
	}
}
