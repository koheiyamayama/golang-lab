package main

type ServiceCer interface {
	DoSomethingC()
}
type ServiceC struct{}

func (s *ServiceC) DoSomethingC() {
	// Do something
}
