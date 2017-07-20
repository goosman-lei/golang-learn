package main

import "fmt"

type StockPosition struct {
    ticker string
    sharePrice float32
    count float32
}

/* method to determine the value of a stock position */
func (s StockPosition) getValue() float32 {
    return s.sharePrice * s.count
}

type Car struct {
    make string
    model string
    price float32
}

/* method to determine the value of a car */
func (c Car) getValue() float32 {
    return c.price
}

/* contract that defines different things that have value */
type Valuable interface {
    getValue() float32
}

/* anything that satisfies the "Valuable" interface is accepted */
func showValue(asset Valuable) {
    fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
    var o Valuable = StockPosition{"Google", 577.20, 4}
    showValue(o)

    o = Car{"BMW", "M3", 66500}
    showValue(o)
}