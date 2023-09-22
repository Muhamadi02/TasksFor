package tasks

import "fmt"

func Task39() {

	var a, b int
	fmt.Println("Введите числа A и B (A<B)")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка", err.Error())
		return
	}

	for i := a; i <= b; i++ {

		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", i)
		}
	}
}

func Task40() {
	var a, b, k int
	fmt.Println("Введите числа A и B (A<B)")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка", err.Error())
		return
	}

	for i := a; i <= b; i++ {
		k++
		for j := 1; j <= k; j++ {
			fmt.Printf("%v ", i)
		}
	}
}
