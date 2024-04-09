package main

type Service struct {
	*ServiceA
	*ServiceB
	*ServiceC
}

type OneService interface {
	ServiceAer
	ServiceBer
	ServiceCer
}

func NewOneService() OneService {
	return &Service{
		&ServiceA{},
		&ServiceB{},
		&ServiceC{},
	}
}

type AnotherService interface {
	ServiceAer
	ServiceBer
}

func NewAnotherService() AnotherService {
	return &Service{
		&ServiceA{},
		&ServiceB{},
		nil,
	}
}

func main() {
	oneService := NewOneService()
	oneService.DoSomethingA()
	oneService.DoSomethingB()
	oneService.DoSomethingC()

	anotherService := NewAnotherService()
	anotherService.DoSomethingA()
	anotherService.DoSomethingB()
	// anotherService.DoSomethingC() この呼び出しはエラーになる
}
