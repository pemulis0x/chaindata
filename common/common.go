package common

import "net/http"
import "encoding/json"
import "io"
import "strconv"
import "strings"

/* 
the first two functions will get called in most..
other api implementations. the latter few are mostly
for quick fix testing and miscellaneous functionality
and probably shouldn't be used in production
because they really aren't very safe.
*/

// hit a url and get a body message in bytes..
// .. and a success/error code (200, 404, etc..)
func GetHttpBody(_url string) ([]byte, int) {
    resp, err := http.Get(_url)
    if err != nil{
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    return body, resp.StatusCode
}


// prefix should look like: "https://xyz.com/api/vx/a/b/c"
// params should look like: "variable=value"
func FormURL(_prefix string, _params ...string) (string) {
    params := strings.Join(_params, "&")
    return _prefix + "?" + params
}

// String to Float converter
func StringToFloat(_s string) (float64) {
    res, err := strconv.ParseFloat(_s, 64)
    if err != nil {
        panic(err)
    }
    return res
}



// bytes -> interface map
func json_to_map(_byt []byte) (map[string] interface{}) {
    var dat map[string] interface{}
    json.Unmarshal(_byt, &dat)
    return dat
}
// url -> interface map
func call_untrusted_api(_url string) (map[string] interface{}) {
    body, _ := GetHttpBody(_url)
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
func singlePriceApi(_query string, _url string) (float64) {
    interf := call_untrusted_api(_url)
    res := float_query(interf, _query)
    return res
}

