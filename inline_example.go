package main

import "fmt"

type Person struct {
	Age int
}

//=====smell===========================================================
func getPrice(person Person) string {
	if isChildren(person) {
		return "free"
	}
	return "$15"
}
func isChildren(person Person) bool {
	return person.Age < 12
}

//=====refactor===========================================================
func getPriceRefactor(person Person) string {
	if person.Age < 12 {
		return "free"
	}
	return "$15"
}
func main() {
	mark := Person{Age: 12}
	ai_chan := Person{Age: 11}
	result := getPrice(mark)
	fmt.Println(result)
	result = getPrice(ai_chan)
	fmt.Println(result)
	result = getPriceRefactor(mark)
	fmt.Println(result)
	result = getPriceRefactor(ai_chan)
	fmt.Println(result)
}
