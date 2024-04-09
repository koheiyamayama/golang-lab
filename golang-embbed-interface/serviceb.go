package main

type ServiceBer interface {
	DoSomethingB()
}
type ServiceB struct{}

func (s *ServiceB) DoSomethingB() {
	// Do something
}
