package pubsub

// subscribes to a topic and returns the channel to listen on
func (e *EventBus) Subscribe(topic string) *EventChannel {
	channel := make(EventChannel)
	if subscribers, found := e.subscribers[topic]; found {

	} else {
		e.subscribers[topic] = append(EventChannels{}, channel)
	}
	return &channel
}
