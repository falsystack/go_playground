package chap14_context

type Store interface {
	Fetch() string
	Cancel()
}
