package observable

type Operation string

const (
	// ObjectSet sets or replaces the contents of the object; value is the new object
	ObjectSet Operation = "oset"

	// ListSet sets or replaces the contents of the list; value is the contents of the list
	ListSet Operation = "lset"
	// ListAppend appends a single element to the end of the list; value is the single element
	ListAppend Operation = "lappend"
	// ListConcat appends multiple elements to the end of the list; value is the list of elements to concat
	ListConcat Operation = "lconcat"
)

type Event struct {
	// Operation denotes what happened to the object
	Operation Operation `json:"op"`
	// Value is the data supporting the operation
	Value interface{} `json:"v"`
}

type Observer interface {
	Observe(event Event)

	Equals(other Observer) bool
}

type ObserverFunc func(event Event)

type observerImpl struct {
	key      string
	observer ObserverFunc
}

func NewObserver(key string, observer ObserverFunc) Observer {
	return &observerImpl{
		key:      key,
		observer: observer,
	}
}

func (o *observerImpl) Equals(other Observer) bool {
	if otherImpl, ok := other.(*observerImpl); ok {
		return o.key == otherImpl.key
	}
	return false
}

func (o *observerImpl) Observe(event Event) {
	o.observer(event)
}
