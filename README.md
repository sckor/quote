# Quote
A golang based utility library modelled after the sql package for retrieving simple stock quotes

# Overview
Probably the easiest way to get going is to use go get to download the Yahoo finance quote driver which also includes a sample
executable that shows how to make use of the quote utility.

The yahoo driver can be found here: https://github.com/sckor/yahoo

The sample code is located at: github.com/sckor/yahoo/yahoo-quote/yahoo-quote.go

# Usage

Import the quote package and the driver package (in this case, it shows the Yahoo finance driver)

```go
import (
	"github.com/sckor/quote"
	_ "github.com/sckor/yahoo"
)
```

Get access to the Quote Source by using Open and pass in the name of the driver and any additional driver args

```go
qs, err := quote.Open("yahoo", "")
```

You can then retrieve a list of quotes by using the quote Retrieve function and supplying a list of tickers
```go
tickers := []string{"MSFT", "AAPL"}
q, err := quote.Retrieve(qs, tickers)
```
