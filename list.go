package observable

import "sync"

type List struct {
	lock      sync.Mutex
	observers []ListObserver

	list []interface{}
}

type ListOperation string

const (
	// ListSet sets or replaces the contents of the list
	ListSet ListOperation = "set"
	// ListAppend appends a single element to the end of the list
	ListAppend ListOperation = "append"
	// ListConcat appends multiple elements to the end of the list
	ListConcat ListOperation = "concat"
)

type ListEvent struct {
	// Operation denotes what happened to the list
	Operation ListOperation `json:"op"`
	// Elements is the data supporting the operation
	Elements []interface{} `json:"e"`
}

type ListObserverFunc func(event ListEvent)

type ListObserverImpl struct {
	key      string
	observer ListObserverFunc
}

func NewListObserver(key string, observer ListObserverFunc) ListObserver {
	return &ListObserverImpl{
		key:      key,
		observer: observer,
	}
}

func (o *ListObserverImpl) Equals(other ListObserver) bool {
	if otherImpl, ok := other.(*ListObserverImpl); ok {
		return o.key == otherImpl.key
	}
	return false
}

func (o *ListObserverImpl) Observe(event ListEvent) {
	o.observer(event)
}

func NewList() *List {
	return &List{}
}

func (o *List) Subscribe(observer ListObserver) {
	if observer == nil {
		return
	}

	defer o.lock.Unlock()
	o.lock.Lock()

	// make sure only one instance is subscribed:
	o.unsubscribe(observer)
	o.observers = append(o.observers, observer)

	// lists publish the entire list contents on first subscribe:
	observer.Observe(ListEvent{
		Operation: ListSet,
		Elements:  o.list,
	})
}

func (o *List) Unsubscribe(observer ListObserver) {
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

func (o *List) unsubscribe(observer ListObserver) {
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
		observer.Observe(ListEvent{
			Operation: ListSet,
			Elements:  newList,
		})
	}
}

func (o *List) Append(newElement interface{}) {
	defer o.lock.Unlock()
	o.lock.Lock()

	o.list = append(o.list, newElement)
	newElements := []interface{}{newElement}
	for _, observer := range o.observers {
		observer.Observe(ListEvent{
			Operation: ListAppend,
			Elements:  newElements,
		})
	}
}

func (o *List) Concat(newElements ...interface{}) {
	defer o.lock.Unlock()
	o.lock.Lock()

	o.list = append(o.list, newElements)
	for _, observer := range o.observers {
		observer.Observe(ListEvent{
			Operation: ListConcat,
			Elements:  newElements,
		})
	}
}
