package main

import "fmt"
import "github.com/bcl4rk/chaindata/binance"
import "github.com/fatih/color"

const TESTPAIR string = "BTCUSDT"

func main(){
    fmt.Println("Welcome to bcl4rk/chaindata...")
    testrun()
}

func testrun(){
    fmt.Println("running tests...\n")
    binanceStatus()
}

func test_endpoint(s string, n float64) {
    fmt.Printf("%s: ", s)
    color.Cyan("... OK\n")
}

func binanceStatus(){
    color.White("BINANCE:\n")

    //avg price
    //p := binance.AveragePrice(TESTPAIR)
    //fmt.Printf("5m %v AvgPrice:\t\t%v\n",TESTPAIR,p)
    //fmt.Printf("%d", test_endpoint(binance.AveragePrice, TESTPAIR))
    //color.Green("test\n")
    test_endpoint("avg. price", binance.AveragePrice(TESTPAIR))
}
