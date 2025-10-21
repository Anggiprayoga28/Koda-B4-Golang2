package main

import "fmt"

func main() {
	scores := []int{50, 75, 20, 32, 66, 90}

	index66 := 0
	for i := 0; i < len(scores); i++ {
		if scores[i] == 66 {
			index66 = i
			break
		}
	}

	scoresBaru := []int{}
	for i := 0; i < index66; i++ {
		scoresBaru = append(scoresBaru, scores[i])
	}
	scoresBaru = append(scoresBaru, 88)
	for i := index66; i < len(scores); i++ {
		scoresBaru = append(scoresBaru, scores[i])
	}

	for i := 0; i < len(scoresBaru); i++ {
		fmt.Println(scoresBaru[i])
	}
}
