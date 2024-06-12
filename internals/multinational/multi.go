package multinational

import "fmt"

type Multinational struct{}

func (i *Multinational) SayHello() {
	fmt.Println("Hello!")
}

func (i *Multinational) SayHola() {
	fmt.Println("Hola")
}

func (i *Multinational) SayBonjour() {
	fmt.Println("Bonjour")
}
