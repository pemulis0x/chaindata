package binance

import "net/http"
import "encoding/json"
import "io"
import "strconv"
import "strings"


const BINANCE_BASE_URL string = "https://api.binance.com/api/v3/"
const BINANCE_AVG_PRICE_URL string = "avgPrice"


// real one
func AveragePrice(_symbol string) (float64) {
    prefix := BINANCE_BASE_URL + BINANCE_AVG_PRICE_URL
    params := "symbol=" + _symbol
    url := FormURL(prefix, params)

    query := "price"
    price := singleprice_api(query, url)
    return price
}


// base + param slice -> full url
func FormURL(_prefix string, _params ...string) (string) {
    params := strings.Join(_params, "&")
    return _prefix + "?" + params
}

// url -> bytes, error code
func get_http_body(_url string) ([]byte, int) {
    resp, err := http.Get(_url)
    if err != nil{
        panic(err)
    }

    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    return body, resp.StatusCode
}

// bytes -> interface map
func json_to_map(_byt []byte) (map[string] interface{}) {
    var dat map[string] interface{}
    json.Unmarshal(_byt, &dat)
    return dat
}

// url -> interface map
func call_untrusted_api(_url string) (map[string] interface{}) {
    body, _ := get_http_body(_url)
    return json_to_map(body)
}

// interface map -> float64
func float_query(_interf map[string] interface{}, _query string) (float64) {
    res, err := strconv.ParseFloat(_interf[_query].(string), 64)
    if err != nil {
        panic(err)
    }
    return res
}

// url, query -> float64
func singleprice_api(_query string, _url string) (float64) {
    interf := call_untrusted_api(_url)
    res := float_query(interf, _query)
    return res
}

