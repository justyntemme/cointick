package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"net"
)

type coins struct {
	btc  float32
	ltc  string
	doge float32
}

func getBtc(tick *coins) {
	out, err := exec.Command("coinfetch", "btc", "usd").Output()
	out64, err := strconv.ParseFloat(strings.Trim(string(out), "\n"), 32)
	tick.btc = float32(out64)
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
	out, err := exec.Command("coinfetch", "-a ccc", "doge", "usd").Output()
	out64, err := strconv.ParseFloat(strings.Trim(string(out), "\n"), 32)
	tick.doge = float32(out64)
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
