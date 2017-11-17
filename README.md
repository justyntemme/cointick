# cointick
Cointick is a command line ticker that gets cryptocoin data from the goCoinFetch package.


## Flags

* -freq (int): Polling frequency in seconds
* -c (path): Path to TOML config file. 
* -rotate (bool): when true, cointick will rotate which coin is displayed rather than all at once, it will use the freq var as timer for each


## Config file
Tickers = ["opt1","opt2","opt3"]
