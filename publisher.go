package pubsub

// publishes data to a topic by sending it to all registered subscriber channels
func (e *EventBus) Publish(topic string, data interface{}) {
	if subscribers, found := e.subscribers[topic]; found {
		for _, subscriber := range subscribers {
			subscriber <- Event{Data: data, Topic: topic}
		}
	}
}
