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

// Package driver defines the interfaces that must be implemented to retrieve
// a stock quote as used by package quote
package driver

// StockQuote - The result of a stock quote query.
// The symbol is a string and represents the stock ticker
// The LastTradePrice is the price of the stock
type StockQuote struct {
	Symbol         string
	LastTradePrice float64
}

// Driver is the interface that must be implemented by a quote driver
type Driver interface {
	Open(name string) (Handle, error)
}

// Handle provides the access to the quote service
type Handle interface {
	// Retrieve takes a list of stock tickers and returns a list
	// of stock quotes.
	Retrieve(tickers []string) (q []StockQuote, err error)
}
