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
	"net"
	"os"
	"os/exec"
	"time"
)

type coins struct {
	btc  string
	ltc  string
	doge string
}

func getBtc(tick *coins) {
	conn, _ := net.Dial("tcp", "localhost:8888")
	message := make([]byte, 1024)
	fmt.Fprintf(conn, "-a ccc 1 btc usd")
	_, err := conn.Read(message)
	tick.btc = string(message)
	if err != nil {
		return
	}
}

func getLtc(tick *coins) {
	conn, _ := net.Dial("tcp", "localhost:8888")
	message := make([]byte, 1024)
	fmt.Fprintf(conn, "-a ccc 1 ltc usd")
	_, err := conn.Read(message)
	tick.ltc = string(message)
	if err != nil {
		return
	}
}

func getDogecoin(tick *coins) {
	conn, _ := net.Dial("tcp", "localhost:8888")
	message := make([]byte, 1024)
	fmt.Fprintf(conn, "-a ccc 1 doge usd")
	_, err := conn.Read(message)
	tick.doge = string(message)
	if err != nil {
		return
	}
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
	fmt.Println("DOGE/USD \t ", tick.doge)
	if rotate == "true" {
		time.Sleep(500)
	}
}

func main() {

	rotate := flag.String("rotate", "flase", "OPTIONS: true,false. Rotates coins to only show one on screen at a time")
	flag.Parse()

	tick := new(coins)

	for {
		getBtc(tick)
		getLtc(tick)
		getDogecoin(tick)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		updateDisplay(tick, *rotate)
		time.Sleep(2000)
	}

}
