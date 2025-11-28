package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
    LanguageName() string
    Greet(name string) string
}
// German
type germanGreeter struct {
}

func (g germanGreeter) LanguageName() string {
    return "German"
}

func (g germanGreeter) Greet(name string) string {
    return fmt.Sprintf("Hallo %v!", name)
}

func SayHello(name string, g Greeter) string {
    return fmt.Sprintf("I can speak %v: %v",g.LanguageName(), g.Greet(name))
}

// Italy
type Italian struct {
}

func (g Italian) LanguageName() string {
    return "Italian"
}

func (g Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %v!", name)
}
// Portuguese
type Portuguese struct {
}

func (g Portuguese) LanguageName() string {
    return "Portuguese"
}

func (g Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %v!", name)
}