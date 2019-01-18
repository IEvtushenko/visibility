package person

import "fmt"

var (
	Public = 1
	person = 1
)

type Person struct {
	Id     int
	Name   string
	secret string
}

func (p Person) UpdateSecret(secret string) {
	fmt.Println("p.secret = ", p.secret)
	fmt.Println("secret = ", secret)
	p.secret = secret
	fmt.Println("p.secret = ", p.secret)
	fmt.Println("p = ", p)
}
