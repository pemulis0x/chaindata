package binance

import "github.com/bcl4rk/chaindata/common"
import "encoding/json"

const BINANCE_BASE_URL string = "https://api.binance.com/api/v3/"
const BINANCE_AVG_PRICE_URL string = "avgPrice"

/*  
template:

func SomeData(_parameter string) (probably float64 or int) {
    url := 'someurl'            # build an end url 
    body, _ := api_call(url)    # get a bunch of bytes from api

    var out SomeStruct          # declare buffer struct
    json.Unmarshal(..)          # parse the json

    return out.DataPoint        # do stuff w/ the data or return 
} */


type BinanceResponse struct {
    Price   string      `json:"price"`
    Fake    string      `json:"fake"`
}

// 1
func AveragePrice(_symbol string) (float64) {
    prefix := BINANCE_BASE_URL + BINANCE_AVG_PRICE_URL
    params := "symbol=" + _symbol
    url := common.FormURL(prefix, params)

    body, _ := common.GetHttpBody(url)
    var out BinanceResponse
    json.Unmarshal(body, &out)

    return common.StringToFloat(out.Price)
}
