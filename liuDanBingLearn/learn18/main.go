package main

import "fmt"

type Hero struct {
	Name string
	Val  int
	Age  int
}

func (this *Hero) ShowHero() {
	fmt.Println("hero is ", this)
}

func (this *Hero) GetName() string {
	fmt.Println("Name is ", this.Name)
	return this.Name
}
func (this *Hero) SetName(newName string) {

	this.Name = newName
}

func main() {
	var hero Hero
	hero.Name = "123"
	hero.SetName("hahah")
	fmt.Println(hero.GetName())
}
