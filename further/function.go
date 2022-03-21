package main

import "fmt"

type Receiver struct {
	myName    string
	myAge     int
	sleepTime float64
}

/*這裡就是GO語音中的方法*/
/*完整的格式

func(recv receiver_type) methodName (parameter_list)(return_value_list){

}
*/

func (myself Receiver) myDay() (id int) {
	fmt.Println(myself.myName, myself.myAge, myself.sleepTime)
	return 1
}

/*這裡receiver除了不能是指針類型或接口類型，可以是其他任何類型*/

/*指針的方法*/
type point struct {
	container string
}

type implementationer interface {
	show()
}

/*因為這裡是傳的指針，所以會直接改數據*/
func (Point *point) show() {
	Point.container = "this is point one"
}

/*	var sisterOfDC = Receiver{
		myName:    "sisterDC",
		myAge:     25,
		sleepTime: 8.00,
	}
	var myID = sisterOfDC.myDay()
	println(myID)
var One = point{container: "this is point"}
fmt.Println(One)
One.show()
fmt.Println(One)*/
