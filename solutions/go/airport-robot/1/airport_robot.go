package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() (name string)
	Greet(name string) (greeting string)
}

func SayHello(name string, greeter Greeter) (greeting string) {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct{}

func (i Italian) LanguageName() (name string) {
	return "Italian"
}

func (i Italian) Greet(name string) (greeting string) {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() (name string) {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) (greeting string) {
	return fmt.Sprintf("Olá %s!", name)
}
