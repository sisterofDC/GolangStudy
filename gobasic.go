package main

import "fmt"

/*如果導入的包沒有使用，也會編譯通不過*/
var globalVariable int8 = 123

/*指明的變量*/
func testFunction(numberInt int8) {
	println(numberInt)
}

func testFunctionTwo(arr *[3]int) {
	println((*arr)[0])
}

func main() {
	/*	var inputnumber int
		fmt.Scanln(&inputnumber)*/
	var a int
	var b int
	fmt.Scanln(&a, &b)
	println("a:", a, "b:", b)
	/*如果想要交換兩個變量的值，則可以簡單地使用*/
	a, b = b, a
	println("a:", a, "b:", b)
	var numberInt int8 = 111
	testString := "this is test string"
	/*go 語言中必須使用所以被聲明的變量，但未使用的全局變量是沒有問題的*/
	println(numberInt)
	println("this is golang string" + testString)

	intArray := [3]int{1, 2, 3}

	testFunction(numberInt)
	testFunctionTwo(&intArray)

}
