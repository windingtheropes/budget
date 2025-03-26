package argent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// USING THE CANADA GOV API https://bcd-api-dca-ipa.cbsa-asfc.cloud-nuage.canada.ca/exchange-rate-lambda/exchange-rates

type CBSA_ForeignExchangeRates struct {
	ForeignExchangeRates []CBSA_ExchangeRate `json:"ForeignExchangeRates"`
}
type CBSA_Value struct {
	Value string `json:"Value"`
}

type CBSA_ExchangeRate struct {
		ExchangeRateId int `json:"ExchangeRateId"`
		Rate string `json:"Rate"`
		ExchangeRateEffectiveTimestamp string `json:"ExchangeRateEffectiveTimestamp"`
		ExchangeRateExpiryTimestamp string `json:"ExchangeRateExpiryTimestamp"`
		ExchangeRateSource string `json:"ExchangeRateSource"`
		FromCurrency CBSA_Value `json:"FromCurrency"`
		FromCurrencyCSN int `json:"FromCurrencyCSN"`
		ToCurrency CBSA_Value `json:"ToCurrency"`
		ToCurrencyCSN int `json:"ToCurrencyCSN"`
}

func GetExchangeCAD(from string) (float64, error) {
	r, err := http.Get("https://bcd-api-dca-ipa.cbsa-asfc.cloud-nuage.canada.ca/exchange-rate-lambda/exchange-rates")
	if err != nil {
	   log.Fatalln(err)
	   return 0, err
	}

	var	fex CBSA_ForeignExchangeRates;
	json.NewDecoder(r.Body).Decode(&fex)
	rates := fex.ForeignExchangeRates;
	
	for i := 0; i < len(rates); i++ {
		if rates[i].FromCurrency.Value == from && rates[i].ToCurrency.Value == "CAD" {
			fmt.Printf("%s", rates[i].Rate)
			fl, err := strconv.ParseFloat(rates[i].Rate, 64)
			if err != nil {
				return 0, err
			}
			return fl, nil
		}
	}
	return 0, fmt.Errorf("currency conversion error")
}

func GetCurrencies() ([]string, error) {
	r, err := http.Get("https://bcd-api-dca-ipa.cbsa-asfc.cloud-nuage.canada.ca/exchange-rate-lambda/exchange-rates")
	if err != nil {
	   log.Fatalln(err)
	   return nil, err
	}

	var	fex CBSA_ForeignExchangeRates;
	json.NewDecoder(r.Body).Decode(&fex)
	rates := fex.ForeignExchangeRates;
	
	var values []string;
	for i := 0; i < len(rates); i++ {
		var s = rates[i].FromCurrency.Value
		values = append(values, s)
	}
	return values, nil
}