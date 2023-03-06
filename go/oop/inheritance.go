package oop

import "fmt"

type Address struct {
	City  string
	State string
}

type PersonI struct {
	Name    string
	Age     int
	Address Address // композиция
}

type Employee struct {
	PersonI // встраивание
	Salary  int
}

func main() {
	e := Employee{PersonI: PersonI{Name: "John", Age: 30, Address: Address{City: "New York", State: "NY"}}, Salary: 50000}
	fmt.Println(e.Name, e.Age, e.Address.City, e.Salary)
}

//Здесь мы создаем структуры Address и Person.
//Затем мы создаем структуру Employee, которая включает в себя структуру Person в качестве поля.
//Мы также используем композицию, чтобы добавить поле Address к структуре Person.
//В итоге мы можем создавать экземпляры структуры Employee с доступом к полям из структур Person и Address.
