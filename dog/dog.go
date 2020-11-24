package dog

import "log"

// Dog woof...
type Dog struct{}

// Bark woof
func (d Dog) Bark() {
	log.Println("WOOF!!!")
}
