package functions

import "fmt"

// Также в Go есть возможность определения функций с переменным числом аргументов (varadic functions).
// Вот пример функции sum, которая принимает переменное количество аргументов типа int и возвращает их сумму:
func sum(numbers ...int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	fmt.Println(sum(1, 2, 3))    // выводит "6"
	fmt.Println(sum(4, 5, 6, 7)) // выводит "22"
}
