package main

import (
	"fmt"
	
	"mymodule/mypkg"
	"mymodule/greeting"
)

func main() {
	fmt.Println("main")
	mypkg.Do()
	greeting.Hello()

	user1 := greeting.User {
		Name: "Jisoo",
		Age: 28,
	}
	user1.Greet()
}