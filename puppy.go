package puppy

import (
	"fmt"

	"github.com/andrehdsc/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}
func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

// versions testing
func ver11() {
	fmt.Println("This is version v1.1.0")
}
func ver12() {
	fmt.Println("This is version v1.2.0")
}
func ver13() {
	fmt.Println("This is version v1.3.0")
}
