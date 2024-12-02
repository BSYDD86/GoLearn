package main

import "fmt"

type AnimalIf interface {
	Sleep()
	GetColor()
	GetType() string
}

func ShowAnimal(this AnimalIf) {
	fmt.Println("animal type", this.GetType())
	this.GetColor()
	this.GetType()

}

type Cat struct {
	Color string
}

func (this *Cat) GetColor() string {
	return this.Color
}
func (this *Cat) Sleep() {
	fmt.Println("cat sleep")
}
func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	Color string
}

func (this *Dog) GetColor() {
	fmt.Println("dog color")
	fmt.Println("color is :", this.Color)
}

func (this *Dog) Sleep() {
	fmt.Println("dog sleep")
}

func (this *Dog) GetType() string {
	return "Dog"
}

func main() {
	var animal AnimalIf
	animal = &Dog{Color: "red"}

	animal.Sleep()
	animal.GetColor()
	animal.GetType()

	ShowAnimal(animal)
	//fmt.Println(dog.Color)
	//fmt.Println(dog.GetColor)
	//fmt.Println(dog.GetType())

}
