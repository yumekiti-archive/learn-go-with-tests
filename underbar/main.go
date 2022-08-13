package main

import "fmt"

func Loop(numbers []int) string {
	tmp := []int{}

	for _, n := range numbers {
		tmp = append(tmp, n)
	}

	return fmt.Sprint(tmp)
}

func main() {
	numbers := []int{20, 30, 40}

	fmt.Println(Loop(numbers))
}
