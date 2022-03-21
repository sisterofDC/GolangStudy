package main

import "fmt"

/*在GO中，可以定義一維或者多萬數組，GO語音中數組是一種值的類型，所以可以通過NEW()來進行創建*/

const name = "sisterOfDC"

func ArrayStatementAndUse() {
	var arrayNumber = [5]int{1, 2, 3, 4, 5}
	var arrayString = [2]string{"this is one", "this is two"}
	var newArray = new([10]int) /*這裡是創建的是數組的指針*/
	var basicArray [10]int
	var arrayOne = newArray
	var arrayTwo = basicArray
	newArray[0] = 1
	basicArray[0] = 2
	fmt.Println(newArray, arrayOne, "地址為", &newArray, &arrayOne)
	fmt.Println(basicArray, arrayTwo, "地址為", &basicArray, &arrayTwo)
	/*	&[1 0 0 0 0 0 0 0 0 0] &[1 0 0 0 0 0 0 0 0 0] 地址為 0xc0000ca018 0xc0000ca020
		[2 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] 地址為 &[2 0 0 0 0 0 0 0 0 0] &[0 0 0 0 0 0 0 0 0 0]*/

	fmt.Println(arrayNumber, arrayString)
}

/*和C語音一樣，同樣會編譯沒有使用的函數*/
/*在寫函數或方法，如果參數是數組，參數長度不能過大*/

func sliceStudy() {
	/*make()是可以申請slice的，算是golang中比較特殊的存在吧，具體用法不是很清楚*/
	var quoteArray = [5]int{1, 2, 3, 4, 5}
	var sliceOne = quoteArray[2:5]
	var sliceTwo = make([]int, 5)
	fmt.Println(sliceOne, sliceTwo)
}

func mapStudy() {
	/*字典申請方法*/
	var mapOne = map[int]string{
		1: "one",
		2: "two"}
	/*遍歷輸出*/
	for i, j := range mapOne {
		println(i, j)
	}
}
