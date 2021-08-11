package main

import "fmt"
import "github.com/bcl4rk/chaindata/binance"
import "github.com/bcl4rk/chaindata/opensea"
import "github.com/fatih/color"

const TESTPAIR string = "BTCUSDT"
const TESTNFT string = "0x47e22659d9ae152975e6cbfa2eed5dc8b75ac545"

func main(){
    fmt.Println("Welcome to bcl4rk/chaindata...")
    testrun()
}

func testrun(){
    fmt.Println("running tests...\n")
    binanceStatus()
    openseaStatus()
}

func test_endpoint(s string, n float64) {
    fmt.Printf("%s: ", s)
    color.Cyan("... OK\n")
}

func binanceStatus(){
    color.White("BINANCE:\n")
    test_endpoint("avg. price", binance.AveragePrice(TESTPAIR))
}

func openseaStatus(){
    opensea.GetAsset(TESTNFT, "1")
}
