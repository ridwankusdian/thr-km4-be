package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		newDataSlice := []int{newData}
		return append(newDataSlice, data...)
	} else if position == "down" {
		return append(data, newData)
	} else {
		panic("Invalid position")
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	newData := 6
	position := "up"
	result := AddElement(data, newData, position)
	fmt.Println(result) // Output: [6 1 2 3 4 5]

	position = "down"
	result = AddElement(data, newData, position)
	fmt.Println(result) // Output: [1 2 3 4 5 6]
}
