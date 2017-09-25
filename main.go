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
	"time"

	"github.com/justyntemme/goCoinFetch"
)

type coins struct {
	btc  string
	ltc  string
	doge string
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
		time.Sleep(500)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	fmt.Println("LTC/USD \t ", tick.ltc)
	if rotate == "true" {
		time.Sleep(500)
		cmd := exec.Command("clear")
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
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		updateDisplay(tick, *rotate)
		time.Sleep(2000)
	}

}
