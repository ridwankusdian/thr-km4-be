package main

import "fmt"

func howManyElements(data []any) int {
	return len(data)
}

func main() {
	data1 := []any{1, 2, 3, 4, 5}
	fmt.Println(howManyElements(data1)) // Output: 5

	data2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(howManyElements(data2)) // Output: 10

	data3 := []interface{}{1, 1, 1, 5, 5, 5}
	fmt.Println(howManyElements(data3)) // Output: 6

	data4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(data4)) // Output: 4

	data5 := []interface{}{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(data5)) // Output: 8

	data6 := []interface{}{true, false, true, false, true, false}
	fmt.Println(howManyElements(data6)) // Output: 6
}
