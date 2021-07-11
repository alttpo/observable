package observable

import (
	"sync"
)

type Object struct {
	lock      sync.Mutex
	observers []Observer

	object interface{}
}

func NewObject() *Object {
	return &Object{}
}

func (o *Object) Subscribe(observer Observer) {
	if observer == nil {
		return
	}

	defer o.lock.Unlock()
	o.lock.Lock()

	// make sure only one instance is subscribed:
	o.unsubscribe(observer)
	o.observers = append(o.observers, observer)

	// publish the current state to new subscribers:
	observer.Observe(Event{
		Operation: ObjectSet,
		Value:     o.object,
	})
}

func (o *Object) Unsubscribe(observer Observer) {
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

func (o *Object) unsubscribe(observer Observer) {
	for i := len(o.observers) - 1; i >= 0; i-- {
		if observer.Equals(o.observers[i]) {
			o.observers = append(o.observers[0:i], o.observers[i+1:]...)
		}
	}
}

func (o *Object) Object() interface{} {
	return o.object
}

func (o *Object) Set(object interface{}) {
	defer o.lock.Unlock()
	o.lock.Lock()

	o.object = object
	for _, observer := range o.observers {
		observer.Observe(Event{
			Operation: ObjectSet,
			Value:     object,
		})
	}
}
