package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

// Language interface
type Greeter interface {
	Greet(name string) string
	LanguageName() string
}

// Italian struct
type Italian struct{}

func (i Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

func (i Italian) LanguageName() string {
	return "Italian"
}

// Portuguese struct
type Portuguese struct{}

func (p Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func SayHello(name string, g Greeter) string {
	return "I can speak " + g.LanguageName() + ": " + g.Greet(name)
}

