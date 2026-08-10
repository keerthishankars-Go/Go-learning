package main

// So far we have only been storing the currency name in the map. If we are able to store the symbol of the currency too? This can be achieved by using a map of structs.

// The currency can be represented as a struct containing the fields currency name and currency symbol. This struct value can be stored in the map with a currency code as key .

import (
	"fmt"
)

type currency struct {
	name   string
	symbol string
}

func main() {
	curUSD := currency{
		name:   "US Dollar",
		symbol: "$",
	}
	curGBP := currency{
		name:   "Pound Sterling",
		symbol: "£",
	}
	curEUR := currency{
		name:   "Euro",
		symbol: "€",
	}

	currencyCode := map[string]currency{
		"USD": curUSD,
		"GBP": curGBP,
		"EUR": curEUR,
	}

	for cyCode, cyInfo := range currencyCode {
		fmt.Printf("CurrencyCode: %s, Name: %s, Symbol: %s\n", cyCode, cyInfo.name, cyInfo.symbol)
	}

}

// In the above program, currency struct contains fields name and symbol. We create three currencies curUSD, curGBP and curEUR.

// In line no. 30, we initialize a map with string key and value of type currency with the three currencies we created.

// The map is iterated in line no. 36 and the currency details are printed in the next line. This program will print,

// Currency Code: USD, Name: US Dollar, Symbol: $
// Currency Code: GBP, Name: Pound Sterling, Symbol: £
// Currency Code: EUR, Name: Euro, Symbol: €
