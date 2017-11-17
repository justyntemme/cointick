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
	"os"
	"os/exec"
	"time"

	"github.com/justyntemme/cointick/pflags"

	"github.com/justyntemme/cointick/clear"
	"github.com/justyntemme/cointick/configReader"
	"github.com/justyntemme/goCoinFetch"
)

type flags struct {
	Freq       int
	Rotate     bool
	ConfigPath string
	Cow        bool
}
type config struct {
	tickers []string
}

func init() {
	clear.Init()

}

func main() {
	f := pflags.ParseFlags()

	configReader.ParseConfig(f.ConfigPath)

	if f.Rotate == true {
		for {
			for _, element := range configReader.ReturnTickers() {
				if f.Cow == true {
					cmd := exec.Command("cowsay " + goCoinFetch.GrabTicker(element))
					cmd.Stdout = os.Stdout
					cmd.Run()
				} else {
					print(element + "/USD\n" + goCoinFetch.GrabTicker(element) + "\n")
					time.Sleep(time.Duration(f.Freq) * time.Second)
					clear.ClearScreen()
				}
			}

		}
	}
	for {
		for _, element := range configReader.ReturnTickers() {
			if f.Cow == true {
				cmd := exec.Command("cowsay", goCoinFetch.GrabTicker(element))
				cmd.Stdout = os.Stdout
				cmd.Run()
			} else {
				print(element + "/USD\n" + goCoinFetch.GrabTicker(element) + "\n")
			}
		}
		time.Sleep(time.Duration(f.Freq) * time.Second)
		clear.ClearScreen()
	}

}
