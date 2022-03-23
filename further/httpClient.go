package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const urlAddress = "http://192.168.240.131:8080/"

type Dbplayer struct {
	Id      int
	Name    string
	Itemone int
	Itemtwo int
}

func doReuqest() {
	var playerOne = Dbplayer{
		Id:      10,
		Name:    "sisterofdc",
		Itemone: 2,
		Itemtwo: 2,
	}
	/*这里还差自定义头，差错检测*/
	jsonstr, err := json.Marshal(playerOne)
	if err != nil {
		fmt.Println(err)
	}
	var client = &http.Client{}
	request, err := http.NewRequest("POST", urlAddress, bytes.NewBuffer(jsonstr))
	if err != nil {
		fmt.Println(err)
	}
	response, err2 := client.Do(request)
	if err2 != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
}

func main() {
	doReuqest()
	/*	time.Sleep(time.Second * 5)*/
}
