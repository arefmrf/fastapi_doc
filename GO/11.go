package main

import "fmt"

func calcArea(r float64) float64 {
	return r * r * 3.14
}
func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}
func calcDiameter(r float64) float64 {
	return 2 * r
}

func printResult(radius float64, calcFunction func(float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
}

func getFunction(query int) func(float64) float64 {
	queryToFunc := map[int]func(float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return queryToFunc[query]
}

func main() {
	var radius float64
	var query int
	fmt.Print("Enter radius: ")
	_, err := fmt.Scanf("%f", &radius)
	if err != nil {
		return
	}
	fmt.Print("Enter query: ")
	_, err = fmt.Scanf("%d", &query)
	if err != nil {
		return
	}
	printResult(radius, getFunction(query))
}
