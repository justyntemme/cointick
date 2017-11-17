package configReader

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type config struct {
	Tickers []string
}

var c = new(config)

func ParseConfig(configPath string) {
	if configPath != "" {
		tomlBytes, err := ioutil.ReadFile(configPath)
		if err != nil {
			print("Error:" + err.Error())
		}
		tomlData := string(tomlBytes)

		if _, err := toml.Decode(tomlData, &c); err != nil {
			print("Error:" + err.Error())
		}

	} else {
		c.Tickers = append(c.Tickers, "BTC")
	}
}
func ReturnTickers() []string {
	return c.Tickers
}
