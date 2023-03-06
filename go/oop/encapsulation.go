package oop

import "fmt"

type Person struct {
	name string // приватное поле
	age  int    // приватное поле
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) GetName() string {
	return p.name
}

func main() {
	p := Person{}
	p.SetName("John")
	fmt.Println(p.GetName())
}

//Здесь мы создаем структуру Person с двумя приватными полями: name и age.
//Затем мы создаем два метода - SetName и GetName - для установки и получения значения поля name соответственно.
//Таким образом, мы скрываем приватные поля и предоставляем публичные методы для доступа к ним.
