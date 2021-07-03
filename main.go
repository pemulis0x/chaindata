package main

import "fmt"
import "github.com/bcl4rk/chaindata/binance"

func main(){
    fmt.Println("Welcome to bcl4rk/chaindata")
    p := binance.AveragePrice("BTCUSDT")
    fmt.Println(p)
}
