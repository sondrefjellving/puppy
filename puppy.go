package puppy

import (
	"fmt"
	"strings"

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
		Name: name,
		Race: race,
		Age: age,
	}
	
	fmt.Println(dog.Dog.GetDogCredentials(gunnarTheDog))
	fmt.Printf("%s barks: %s", strings.ToTitle(gunnarTheDog.Name), Barks())
	fmt.Printf("%s jumps: %s", strings.ToTitle(gunnarTheDog.Name), dog.Jump())
}