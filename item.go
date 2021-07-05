package main

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	i.inStock = true
	i.notifyAll()
}

func (i *item) subscribe(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) unsubscribe(o observer) {
	i.observerList = removeFromObserverList(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}
func removeFromObserverList(list []observer, o observer) (result []observer) {
	for i, item := range list {
		if o.getEmail() == item.getEmail() {
			result = append(list[:i], list[i+1:]...)
		}
	}
	return result
}
