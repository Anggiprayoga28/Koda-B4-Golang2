package main

import "fmt"

func main() {
	scores := []int{50, 75, 66, 20, 32, 90}

	scores = append(scores[:3], append([]int{88}, scores[3:]...)...)

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
}
