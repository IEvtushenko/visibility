package main

import (
	"coursera/visibility/person"
	"fmt"
)

func main() {
	p := person.NewPerson(1, "name", "secret")

	secret := person.GetSecret(p)
	fmt.Println("GetSecret", secret)
	p.UpdateSecret("double")

	//secret = person.GetSecret(p)
	//fmt.Println("GetSecret", secret)
}
