package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet() string
}

type German struct {
}

func (g German) LanguageName() string {
    return "German"
}

func (g German) Greet() string {
    return "Hallo"
}



type Italian struct {
}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet() string{
    return "Ciao"
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet() string{
    return "Olá"
}


func SayHello(name string, greeter Greeter) string {
    message := fmt.Sprintf("I can speak %s: %s %s!", greeter.LanguageName(), greeter.Greet(), name)
    return message
}