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

func (c *coins) getBtc() {
	c.btc = goCoinFetch.GrabTicker("btc")
}

func (c *coins) getLtc() {
	c.ltc = goCoinFetch.GrabTicker("LTC")
}

func (c *coins) Tick() {
	c.getBtc()
	c.getLtc()
	fmt.Println("BTC/USD \t ", c.btc)
	fmt.Println("LTC/USD \t ", c.ltc)
}

func clearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
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

	var freq int

	flag.IntVar(&freq, "freq", 10, "Polling frequency in seconds")
	flag.Parse()

	tick := new(coins)

	for {
		tick.Tick()
		time.Sleep(time.Duration(freq) * time.Second)
	}

}
