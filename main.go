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
	"io/ioutil"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/justyntemme/cointick/clear"
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
	tickers := []string{}
	tickersN := []string{}

	flag.StringVar(&configPath, "config", "", "Where to find the config file for multiple tickers")
	flag.IntVar(&freq, "freq", 10, "Polling frequency in seconds")
	flag.BoolVar(&rotate, "rotate", false, "Displays one ticker at a time when set to true")
	flag.Parse()

	if configPath != "" {
		tomlBytes, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Print("Error:" + err.Error())
		}
		tomlData := string(tomlBytes)
		var c config
		if _, err := toml.Decode(tomlData, &c); err != nil {
			fmt.Println("Error:" + err.Error())
		}
		for index, _ := range c.tickers {
			tickersN = append(tickersN, c.tickers[index])
			tickers = append(tickers, goCoinFetch.GrabTicker(c.tickers[index]))

		}
	}
	// Checks if no config path is used, if so defualts to BTC
	if configPath == "" {
		tickersN = append(tickersN, "btc")
		tickers = append(tickers, goCoinFetch.GrabTicker("btc"))
	}
	if rotate == true {
		for {
			for index, _ := range tickersN {
				fmt.Println(tickersN[index] + "/USD\n" + goCoinFetch.GrabTicker(tickersN[index]))
				time.Sleep(time.Duration(freq) * time.Second)
				clear.ClearScreen()
			}

		}
	}
	for {
		for index, _ := range tickers {
			fmt.Println(tickersN[index] + "/USD\n" + goCoinFetch.GrabTicker(tickersN[index]))
		}
		time.Sleep(time.Duration(freq) * time.Second)
		clear.ClearScreen()
	}

}
