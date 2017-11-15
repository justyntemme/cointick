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

func clearScreen() {
	value, ok := clear[runtime.GOOS]
	//runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
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
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

var clear map[string]func()

func main() {

	freq := 10
	rotate := false
	tickers := []string{}
	tickersN := []string{}

	flag.IntVar(&freq, "freq", 10, "Polling frequency in seconds")
	flag.BoolVar(&rotate, "rotate", false, "Displays one ticker at a time when set to true")
	flag.Parse()

	//Hardcoding will add flag functionality soon
	tickersN = append(tickersN, "btc")
	tickersN = append(tickersN, "ltc")
	tickers = append(tickers, goCoinFetch.GrabTicker("btc"))
	tickers = append(tickers, goCoinFetch.GrabTicker("ltc"))

	if rotate == true {
		for {
			for index, _ := range tickers {
				fmt.Println(tickersN[index] + "/USD\n" + tickers[index])
				time.Sleep(time.Duration(freq) * time.Second)
				clearScreen()
			}

		}
	}
	for {
		for index, _ := range tickers {
			fmt.Println(tickersN[index] + "/USD\n" + tickers[index])
		}
		time.Sleep(time.Duration(freq) * time.Second)
		clearScreen()
	}

}
