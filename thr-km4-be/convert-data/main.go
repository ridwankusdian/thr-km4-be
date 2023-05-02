package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	fields := strings.Split(data, ",")
	age, _ := strconv.Atoi(fields[1])
	return User{Name: fields[0], Age: age, Address: fields[2]}
}

func main() {
	data1 := "Edo,20,Jakarta"
	user1 := ConvertData(data1)
	fmt.Printf("User 1: %+v\n", user1)

	data2 := "Budi,30,Bandung"
	user2 := ConvertData(data2)
	fmt.Printf("User 2: %+v\n", user2)
}
