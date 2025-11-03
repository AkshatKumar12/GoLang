package main

import "fmt"

// enumerated types

type OrderStatus string

const (
	Received  OrderStatus = "received" // use iota instead of string for numeric data-types for auto sequencing
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Prepared)
}
