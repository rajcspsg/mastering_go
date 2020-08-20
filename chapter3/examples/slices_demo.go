package examples

import "fmt"
import "sort"

func ExamplesWithSlices() {
	aSlice := []int {1, 2, 3, 4, 5};
	fmt.Println(aSlice)

	integers := make([]int, 2)
	fmt.Println(integers)

	integers = nil
	fmt.Println(integers)
}


func ExamplesWithArraysAndSlices() {
	anArray := [5]int { -1, -2, -3, -4, -5 }
	refAnArray := anArray[:]

	fmt.Println("anArray \t ", anArray)
	fmt.Println("refAnArray \t", refAnArray)

	anArray[4] = -100

	fmt.Println("anArray \t ", anArray)
	fmt.Println("refAnArray \t", refAnArray)

}

func MultiDimensionalSlices() {

	s:= make([]byte, 5)
	fmt.Println(s)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()

	for i := 0; i < len(twoD); i++ {
		for j:= 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i * j)
		}
	}


	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i: ", i, "value:", y)
		}
		fmt.Println("")
	}
}


type aStructure struct {
	person string
	height int
	weight int
}

func SortingSlices() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})
	fmt.Println("0:", mySlice)


	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})

	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})

	fmt.Println(">:", mySlice)
}
