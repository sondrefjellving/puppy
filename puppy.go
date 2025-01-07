package puppy

import (
	"fmt"

	"github.com/sondrefjellving/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func PrintDogCreds(name, race string, age int) {
	gunnarTheDog := dog.Dog{
		Name: "Gunnar",
		Race: "Chihuahua",
		Age: 1,
	}
	
	fmt.Println(dog.Dog.GetDogCredentials(gunnarTheDog))
	fmt.Printf("%s barks: %s", gunnarTheDog.Name, Barks())
	fmt.Printf("%s jumps: %s", gunnarTheDog.Name, dog.Jump())
}