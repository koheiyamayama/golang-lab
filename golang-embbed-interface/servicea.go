package main

type ServiceAer interface {
	DoSomethingA()
}
type ServiceA struct{}

func (s *ServiceA) DoSomethingA() {
	// Do something
}
