package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"net"
)

type coins struct {
	btc  string
	ltc  string
	doge string
}

func getBtc(tick *coins) {
	conn, _ := net.Dial("tcp","localhost:8888")
	message := make([]byte, 1024)
	fmt.Fprintf(conn,"-a ccc 1 btc usd")
	_, err := conn.Read(message)
	tick.btc = string(message)
	if err != nil {
		return
	}
}

func getLtc(tick *coins) {
	conn, _ := net.Dial("tcp","localhost:8888")
	message := make([]byte, 1024)
	fmt.Fprintf(conn,"-a ccc 1 ltc usd")
	_, err := conn.Read(message)
	tick.ltc = string(message)
	if err != nil {
		return
	}
}

func getDogecoin(tick *coins) {
	conn, _ := net.Dial("tcp","localhost:8888")
	message := make([]byte, 1024)
	fmt.Fprintf(conn,"-a ccc 1 doge usd")
	_, err := conn.Read(message)
	tick.doge = string(message)
	if err != nil {
		return
	}
}

func updateDisplay(tick *coins) {
	fmt.Println("BTC/USD \t ", tick.btc)
	fmt.Println("LTC/USD \t ", tick.ltc)
	fmt.Println("DOGE/USD \t ", tick.doge)
}

func main() {
	tick := new(coins)
	for {
		getBtc(tick)
		getLtc(tick)
		getDogecoin(tick)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		updateDisplay(tick)
		time.Sleep(20)
	}
}
