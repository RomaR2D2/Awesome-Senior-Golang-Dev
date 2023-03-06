package oop

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	p := person{Name: "John", Age: 30}
	fmt.Println(p.Name, p.Age)
}

//Здесь мы создаем структуру Person с двумя полями: Name и Age. Затем мы создаем экземпляр этой структуры и выводим его поля.
