package pflags

import "flag"

type flags struct {
	Freq       int
	Rotate     bool
	ConfigPath string
}

func ParseFlags() *flags {
	freq := 10
	rotate := false
	configPath := ""

	flag.StringVar(&configPath, "config", "", "Where to find the config file for multiple tickers")
	flag.IntVar(&freq, "freq", 10, "Polling frequency in seconds")
	flag.BoolVar(&rotate, "rotate", false, "Displays one ticker at a time when set to true")
	flag.Parse()
	f := new(flags)
	f.ConfigPath = configPath
	f.Freq = freq
	f.Rotate = rotate
	return f
}
