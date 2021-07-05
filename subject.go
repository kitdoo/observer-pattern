package main

type subject interface {
	subscribe(Observer observer)
	unsubscribe(Observer observer)
	notifyAll()
}
