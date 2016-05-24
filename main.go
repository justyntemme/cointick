package main

import (
	"os/exec"
	"strconv"
)

type coins struct {
	btc float32
	ltc float32
	dogecoin float32
}

func getBtc() float32 {
	var val float32


	return	val
}

func getLtc() float32 {
	var val float32

	return val
}

func dogecoin() float32 {
	var val float32


	return val
}


func main() {
	tick := new(coins)
	out,err := exec.Command("btc","usd").Output()
	out64,err := strconv.ParseFloat(string(out), 32)
	tick.btc = float32(out64)
	if err == nil{
		return
	}else {
		return
	}
}
