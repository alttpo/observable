package observable

import (
	"sync"
)

type Object struct {
	lock      sync.Mutex
	observers []ObjectObserver

	object interface{}
}

type ObjectOperation string

const (
	ObjectSet ObjectOperation = "set"
)

type ObjectEvent struct {
	// Operation denotes what happened to the object
	Operation ObjectOperation `json:"op"`
	// Value is the data supporting the operation
	Value interface{} `json:"v"`
}

type ObjectObserverFunc func(event ObjectEvent)

type objectObserverImpl struct {
	key      string
	observer ObjectObserverFunc
}

func NewObjectObserver(key string, observer ObjectObserverFunc) ObjectObserver {
	return &objectObserverImpl{
		key:      key,
		observer: observer,
	}
}

func (o *objectObserverImpl) Equals(other ObjectObserver) bool {
	if otherImpl, ok := other.(*objectObserverImpl); ok {
		return o.key == otherImpl.key
	}
	return false
}

func (o *objectObserverImpl) Observe(event ObjectEvent) {
	o.observer(event)
}

func NewObject() *Object {
	return &Object{}
}

func (o *Object) Subscribe(observer ObjectObserver) {
	if observer == nil {
		return
	}

	defer o.lock.Unlock()
	o.lock.Lock()

	// make sure only one instance is subscribed:
	o.unsubscribe(observer)
	o.observers = append(o.observers, observer)

	// publish the current state to new subscribers:
	observer.Observe(ObjectEvent{
		Operation: ObjectSet,
		Value:     o.object,
	})
}

func (o *Object) Unsubscribe(observer ObjectObserver) {
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

func (o *Object) unsubscribe(observer ObjectObserver) {
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
		observer.Observe(ObjectEvent{
			Operation: ObjectSet,
			Value:     object,
		})
	}
}
