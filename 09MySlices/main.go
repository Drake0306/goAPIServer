package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"a", "t", "p"} // syntax one
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "m", "b")
	fmt.Println(fruitList)

	// assign a start and end value of slice as 1 will skip 0 index and 3 means it will skip everything from 3 to more and will print values in middle
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4) // syntax two
	highScores[0] = 234
	highScores[1] = 935
	highScores[2] = 436
	highScores[3] = 537
	// highScores[4] = 737

	highScores = append(highScores, 737, 666, 352)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index

	fmt.Println("Remove a value from slice based on index")
	var cources = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(cources)

	var index int = 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}
