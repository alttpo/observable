package observable

type ObjectObserver interface {
	Observe(event ObjectEvent)

	Equals(other ObjectObserver) bool
}

type ListObserver interface {
	Observe(event ListEvent)

	Equals(other ListObserver) bool
}
