package observer

type eventManager struct {
	subscribers []EventListener
}

func (e *eventManager) Subscribe(listener EventListener) {
	e.subscribers = append(e.subscribers, listener)
}

func (e *eventManager) Unsubscribe(listener EventListener) {
	for i, subscriber := range e.subscribers {
		if subscriber == listener {
			e.subscribers = append(e.subscribers[:i], e.subscribers[i+1:]...)
			break
		}
	}
}

func (e *eventManager) NotifySubscribers() {
	for _, subscriber := range e.subscribers {
		subscriber.Update()
	}
}

func NewEventManager() *eventManager {
	return &eventManager{
		subscribers: make([]EventListener, 0),
	}
}
