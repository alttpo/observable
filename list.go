package observable

import "sync"

type List struct {
	lock      sync.Mutex
	observers []Observer

	list []interface{}
}

func NewList() *List {
	return &List{}
}

func (o *List) Subscribe(observer Observer) {
	if observer == nil {
		return
	}

	defer o.lock.Unlock()
	o.lock.Lock()

	// make sure only one instance is subscribed:
	o.unsubscribe(observer)
	o.observers = append(o.observers, observer)

	// lists publish the entire list contents on first subscribe:
	observer.Observe(Event{
		Operation: ListSet,
		Value:     o.list,
	})
}

func (o *List) Unsubscribe(observer Observer) {
	if observer == nil {
		return
	}

	defer o.lock.Unlock()
	o.lock.Lock()

	if o.observers == nil {
		return
	}

	o.unsubscribe(observer)
}

func (o *List) unsubscribe(observer Observer) {
	for i := len(o.observers) - 1; i >= 0; i-- {
		if observer.Equals(o.observers[i]) {
			o.observers = append(o.observers[0:i], o.observers[i+1:]...)
		}
	}
}

func (o *List) List() []interface{} {
	return o.list
}

func (o *List) Set(newList ...interface{}) {
	defer o.lock.Unlock()
	o.lock.Lock()

	o.list = newList
	for _, observer := range o.observers {
		observer.Observe(Event{
			Operation: ListSet,
			Value:     newList,
		})
	}
}

func (o *List) Append(newElement interface{}) {
	defer o.lock.Unlock()
	o.lock.Lock()

	o.list = append(o.list, newElement)
	for _, observer := range o.observers {
		observer.Observe(Event{
			Operation: ListAppend,
			Value:     newElement,
		})
	}
}

func (o *List) Concat(newElements ...interface{}) {
	defer o.lock.Unlock()
	o.lock.Lock()

	o.list = append(o.list, newElements)
	for _, observer := range o.observers {
		observer.Observe(Event{
			Operation: ListConcat,
			Value:     newElements,
		})
	}
}
