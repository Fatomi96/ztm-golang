package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// j := 10
	// for j < 20 {
	// 	fmt.Println(j)
	// 	j++
	// }

	// k := 20

	// for {
	// 	fmt.Println("Infinite loop for :", k)
	// 	k++
	// 	if k == 30 {
	// 		break
	// 	}
	// }

	sum := 0
	fmt.Println("sum is", sum)

	for i := 1; i < 10; i++ {
		sum += i
		fmt.Println("sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement. sum is", sum)
	}
}
