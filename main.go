package main

import "fmt"
import "github.com/bcl4rk/chaindata/binance"

const TESTPAIR string = "BTCUSDT"

func main(){
    fmt.Println("Welcome to bcl4rk/chaindata...")
    testrun()
}

func testrun(){
    fmt.Println("running tests...\n")
    binanceStatus()
}

func binanceStatus(){
    fmt.Println("querying binance API endpoints..")

    //avg price
    p := binance.AveragePrice(TESTPAIR)
    fmt.Printf("Binance 5m %v AvgPrice:\t\t%v\n",TESTPAIR,p)
}
