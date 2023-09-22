package main

import (
	"main/observer"
	"main/user"
)

func main() {
	eventManager := observer.NewEventManager()
	user := user.New()
	eventManager.Subscribe(user)
	eventManager.NotifySubscribers()
	eventManager.Unsubscribe(user)
}
