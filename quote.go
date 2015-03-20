// Copyright (c) 2015 Sean Kormilo. All Rights Reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package Quote provides a generic interface around retrieving stock quotes

// The quote package must be used in conjuncition with a quote driver
package quote

import (
	"fmt"
	"github.com/sckor/quote/driver"
	"sort"
)

var drivers = make(map[string]driver.Driver)

// Register makes a stock quote driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, driver driver.Driver) {
	if driver == nil {
		panic("quote: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("quote: Register called twice for driver " + name)
	}
	drivers[name] = driver
}

func unregisterAllDrivers() {
	// For tests.
	drivers = make(map[string]driver.Driver)
}

// Drivers returns a sorted list of the names of the registered drivers.
func Drivers() []string {
	var list []string
	for name := range drivers {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

// QuoteSource is used as the handle used to access the quote source
// to retrieve quotes
type QuoteSource struct {
	driver     driver.Driver
	handle     driver.Handle
	sourceName string
}

func Open(driverName, sourceName string) (*QuoteSource, error) {
	driveri, ok := drivers[driverName]

	if !ok {
		return nil, fmt.Errorf("quote: unknown driver %q (forgotten import?)", driverName)
	}

	handle, err := driveri.Open(sourceName)

	if err != nil {
		return nil, err
	}

	qs := &QuoteSource{
		driver:     driveri,
		sourceName: sourceName,
		handle:     handle,
	}

	return qs, nil
}

func Retrieve(qs *QuoteSource, tickers []string) (q []driver.StockQuote, err error) {
	q, err = qs.handle.Retrieve(tickers)
	return
}
