package main

import "fmt"

func updateName(x *string) {
	*x = "Webber"
}

func main() {

	var name = "tifa"
	var m = &name
	fmt.Println("memory addr:", &name)
	fmt.Println("memory addr:", m)
	fmt.Println("value of memory addr:", *m)

	updateName(m)
	fmt.Println(name)

}
