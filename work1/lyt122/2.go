package main

import "fmt"

func getAppleNumbers(appleHeights []int, maxHeight int, deskHeight int) int {
	var count int
	for _, height := range appleHeights {
		if height <= maxHeight+deskHeight {
			count++
		}
	}
	return count
}
func main() {
	var appleHeights [10]int
	var maxHeight int
	for i := 0; i < 10; i++ {
		fmt.Scan(&appleHeights[i])
	}
	fmt.Scan(&maxHeight)
	fmt.Println(getAppleNumbers(appleHeights[:], maxHeight, 30))
}
