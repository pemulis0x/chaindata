
package opensea

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetAsset(contract string, id string) {
	url := "https://api.opensea.io/api/v1/asset/" + contract + "/" + id + "/"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
