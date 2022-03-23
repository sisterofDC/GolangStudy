package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"
)

/*這裡鏈接遠端的地址*/

const IPAddress = "192.168.240.129:999"
const netProtocol = "tcp"
const get = "GET"

func sendMessage(conn *net.TCPConn) {
	/*同樣申請一個空間*/
	/*	var reader = bufio.NewReader(conn)*/
	var message = []byte(conn.LocalAddr().String() + "this is client \n")
	_, err := conn.Write(message)
	fmt.Println("this is message send")
	if err != nil {
		return
	}
}

func reciveMessage(conn *net.TCPConn) {
	var reader = bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(message)
	}
}

func socketClient() {
	var tcpAddr *net.TCPAddr
	tcpAddr, Connerr := net.ResolveTCPAddr(netProtocol, IPAddress)
	/*鏈接遠端服務器*/
	conn, err := net.DialTCP(netProtocol, nil, tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tcpAddr, Connerr, conn, err)
	go sendMessage(conn)
	go reciveMessage(conn)
	time.Sleep(time.Second * 10)
}

func DoRequest() {
	var client = &http.Client{
		Transport:     nil, /*端口*/
		CheckRedirect: nil, /**/
		Jar:           nil, /**/
		Timeout:       0,   /*超時*/
	}
	request, err := http.NewRequest(get, "", nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("", "")
	response, err2 := client.Do(request)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)

}

/*問題一，看Ping的通不,好的根本Ping通,這裡鏈接起來了*/
