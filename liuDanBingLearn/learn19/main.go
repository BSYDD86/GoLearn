package main

import "fmt"

type Man struct {
	name string
	age  int
}

func (this *Man) Eat() {
	fmt.Println("man is eating")
	fmt.Println(this.age)
}

type SuperMan struct {
	Man
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("superman is eating")
	fmt.Println(this.age)

}

func main() {
	superMan := Man{age: 1, name: "haha"}
	fmt.Println(superMan.name)
	fmt.Println("-----------")
	fmt.Println(superMan.age)
	superMan.Eat()

}
