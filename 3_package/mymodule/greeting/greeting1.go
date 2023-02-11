package greeting

import "fmt"

type User struct {
	Name string
	Age int
}

func (user User) Greet() {
	greeting := "Hello! I'm " + user.Name + "I'm" + fmt.Sprint(user.Age) + "years old."
	fmt.Println(greeting)
}