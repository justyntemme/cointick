/*

Copyright (C) 2016 Justyn Temme
This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/justyntemme/goCoinFetch"
)

var clear map[string]func()

type coins struct {
	btc string
	ltc string
}

func getBtc(tick *coins) {
	tick.btc = goCoinFetch.GrabTicker("btc")
}

func getLtc(tick *coins) {
	tick.ltc = goCoinFetch.GrabTicker("LTC")
}

func updateDisplay(tick *coins, rotate string) {
	fmt.Println("BTC/USD \t ", tick.btc)
	if rotate == "true" {
		time.Sleep(5000 * time.Millisecond)
		clearScreen()
	}

	fmt.Println("LTC/USD \t ", tick.ltc)
}

func clearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func init() {

	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func main() {

	rotate := flag.String("rotate", "flase", "OPTIONS: true,false. Rotates coins to only show one on screen at a time")
	flag.Parse()

	tick := new(coins)

	for {
		getBtc(tick)
		getLtc(tick)
		clearScreen()
		updateDisplay(tick, *rotate)
		time.Sleep(10000 * time.Millisecond)
	}

}
