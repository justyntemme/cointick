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
	"time"

	"github.com/justyntemme/cointick/clear"
	"github.com/justyntemme/cointick/configReader"
	"github.com/justyntemme/goCoinFetch"
)

type config struct {
	tickers []string
}

func init() {
	clear.Init()

}

func main() {
	/* IDEAL SCOPE

	parseFlags()
	parseConfig()




	*/

	freq := 10
	rotate := false
	configPath := ""

	flag.StringVar(&configPath, "config", "", "Where to find the config file for multiple tickers")
	flag.IntVar(&freq, "freq", 10, "Polling frequency in seconds")
	flag.BoolVar(&rotate, "rotate", false, "Displays one ticker at a time when set to true")
	flag.Parse()

	configReader.ParseConfig(configPath)

	if rotate == true {
		for {
			for _, element := range configReader.ReturnTickers() {
				fmt.Println(element + "/USD\n" + goCoinFetch.GrabTicker(element))
				time.Sleep(time.Duration(freq) * time.Second)
				clear.ClearScreen()
			}

		}
	}
	for {
		for _, element := range configReader.ReturnTickers() {
			fmt.Println(element + "/USD\n" + goCoinFetch.GrabTicker(element))
		}
		time.Sleep(time.Duration(freq) * time.Second)
		clear.ClearScreen()
	}

}
