package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	user := User{
		Name: "周杰伦",
		Age:  26,
	}

	fmt.Printf("%%\n")
	fmt.Printf("%b\n", user.Age)
	fmt.Printf("%c\n", 64)
	fmt.Printf("%x\n", user.Age)
	fmt.Printf("%X\n", user.Age)
	fmt.Printf("%c\n", user.Name[0])
	fmt.Printf("%U\n", user.Name[0])
	fmt.Printf("%d\n", user.Age)
	fmt.Printf("%t\n", user.Age > 20)
	fmt.Printf("%e\n", 3.1415926)
	fmt.Printf("%20.3e\n", 3.1415926)
	fmt.Printf("%E\n", 3.1415926)
	fmt.Printf("%f\n", 3.1415926)
	fmt.Printf("%o\n", user.Age)
	fmt.Printf("%p\n", []int{1})
	fmt.Printf("%s\n", user.Name)
	fmt.Printf("%q\n", user.Name)
	fmt.Printf("%T\n", user.Name)
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%#v\n", user.Age)
	fmt.Printf("%#v\n", user.Name)
}
