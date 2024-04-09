package main

type Service struct {
	A *ServiceA
	B *ServiceB
	C *ServiceC
}

type ServiceA struct{}

type ServiceB struct{}

type ServiceC struct{}

func NewOneService() *Service {
	return &Service{
		A: &ServiceA{},
		B: &ServiceB{},
		C: &ServiceC{},
	}
}

func NewAnotherService() *Service {
	return &Service{
		A: &ServiceA{},
		B: &ServiceB{},
		C: nil,
	}
}
