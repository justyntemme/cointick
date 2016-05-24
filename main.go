package main

import (
	"os/exec"
	"os"
	"strconv"
	"strings"
	"fmt"
	"time"
)

type coins struct {
	btc float32
	ltc float32
	doge float32
}

func getBtc(tick *coins)  {
	out,err :=exec.Command("coinfetch","btc", "usd").Output()
	out64,err := strconv.ParseFloat(strings.Trim(string(out), "\n"), 32)
	tick.btc = float32(out64)
	if err != nil {
		return
	}
}

func getLtc(tick *coins)  {
	out,err :=exec.Command("coinfetch", "ltc", "usd").Output()
	out64,err := strconv.ParseFloat(strings.Trim(string(out),"\n"), 32)
	tick.ltc = float32(out64)
	if err != nil {
		return
	}
}

func getDogecoin(tick *coins) {
	out,err :=exec.Command("coinfetch", "-a bter", "doge", "usd").Output()
	out64,err := strconv.ParseFloat(strings.Trim(string(out), "\n"), 32)
	tick.doge = float32(out64)
	if err != nil {
		return
	}
}

func updateDisplay(tick *coins) {
	fmt.Println("BTC/USD \t ",tick.btc)
	fmt.Println("LTC/USD \t ",tick.ltc)
	fmt.Println("DOGE/USD \t ",tick.doge)
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
