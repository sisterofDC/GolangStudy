package main

import "fmt"

/*type 關鍵字在GO語言中很重要，定義結構體，接口，類別*/
/*這裡有個重點，go中沒有隱式類型轉換，這裡就說明類型*/

func typeStudy() {
	type Able struct {
		name string
		time float64
	}
	/*感覺還是和C語音差不多*/
	var one = Able{
		name: "one",
		time: 1.2,
	}
	fmt.Println(one)
}

/*函數
func 函數名（參數）（返回參數）｛

return
｝
*/

/*在go語音中沒有函數重載，所以方法都是單一的*/

func testOne(varOne int, varTwo string) (backOne int, backTwo string) {
	var One = varOne + 10
	var Two = varTwo + "this is add"
	return One, Two
}

func functionStudy() {
	type backStruct struct {
		One int
		Two string
	}
	var anonymousFun = func(varOne int) int {
		return varOne + varOne
	}
	var back backStruct
	back.One, back.Two = testOne(1, "one")
	back.One = anonymousFun(2)
	fmt.Println(back)
	/*	{4 one this is add}*/
}

/*內置函數make()和new()都和內存分配有關*/
/*結構體和C語音中沒有什麼區別*/

/*同一個接口可以被多哥類型實現，一個類型也可以實現多哥接口*/
/**/
type playerOne struct {
	name string
}

type playerTwo struct {
	name  string
	speed float64
}

type playerThree struct {
	name      string
	speed     float64
	transform []int
}

type conditioner interface {
	setname() string
	settransform() []int
	setspeed() float64
}

func (player playerOne) setname() {
	fmt.Println(player.name)
}

/*一個類型可以實現所有接口函數,當然go已經可以支持泛型了，就不討論這麼多了*/
/*剩下的就是類型判斷，和類型轉換*/
func (player playerTwo) setspeed() {
	fmt.Println(player.speed)
}
func (player playerTwo) setname() {
	fmt.Println(player.name)
}

func (player playerThree) settransform() {
	fmt.Println(player.transform)
}

func implementInterface() {

}
func interfaceStudy() {

}

/*	typeStudy()*/
/*	functionStudy()*/
/*	var myplay = playerOne{name: "sisterOFDC"}
	myplay.setname()

	var mySecondPlay = playerTwo{
		name:  "two",
		speed: 2,
	}
	mySecondPlay.setname()
	mySecondPlay.setspeed()
*/
