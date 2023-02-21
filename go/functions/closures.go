package functions

import "fmt"

// Замыкания (closures) позволяют сохранять состояние между вызовами функции.
// Вот пример функции counter, которая возвращает другую функцию, которая в свою очередь увеличивает счетчик при каждом вызове:
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	c := counter()
	fmt.Println(c()) // выводит "1"
	fmt.Println(c()) // выводит "2"
	fmt.Println(c()) // выводит "3"
}
