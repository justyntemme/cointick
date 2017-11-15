package config

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type config struct {
	tickers []string
}
var c := new(config)
func ParseConfig(configPath string) {
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
		//	tickersN = append(tickersN, c.tickers[index])
		//	tickers = append(tickers, goCoinFetch.GrabTicker(c.tickers[index]))

	}

}

func ReturnTickers() []string {


}
