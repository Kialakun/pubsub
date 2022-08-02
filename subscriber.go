package pubsub

// subscribes to a topic and returns the channel to listen on
func (e *EventBus) Subscribe(topic string) *EventChannel {
	// @todo: implement mutex to safely read/write to EventBus
	channel := make(EventChannel)
	if _, found := e.subscribers[topic]; found {
		e.subscribers[topic] = append(e.subscribers[topic], channel)
	} else {
		e.subscribers[topic] = append(EventChannels{}, channel)
	}
	return &channel
}
