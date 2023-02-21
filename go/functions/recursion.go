package functions

import "fmt"

// Рекурсивные функции в Go могут быть использованы для решения различных задач, например, для обхода деревьев или поиска в глубину.
// Вот пример функции factorial, которая вычисляет факториал числа с использованием рекурсии:
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5)) // выводит "120"
}
