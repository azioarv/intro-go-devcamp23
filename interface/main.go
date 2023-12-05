package main

import (
	"fmt"
)

type Address struct {
	UserID  int64
	Address string
	ZipCode string
}

type Response struct{}

type Logistic interface {
	CreateShipment(address Address) Response
}

type GoSend struct {
	URL string
}

type GTL struct {
	URL string
}

func (g *GoSend) CreateShipment(address Address) Response {
	// api call process
	fmt.Println("Create Shipment for Go-Send For UserID ", address.UserID)
	return Response{}
}
func (gtl *GTL) CreateShipment(address Address) Response {
	// api call process
	fmt.Println("Create Shipment for GTL For UserID ", address.UserID)
	return Response{}
}

func main() {
	gosend := GoSend{
		URL: "https://api.gojek.com",
	}

	firstAddress := Address{
		UserID:  123,
		Address: "Tokopedia Tower Lt 52",
		ZipCode: "12950",
	}
	generateShipment(&gosend, firstAddress)

	gtl := GTL{
		URL: "https://api.gtl.id",
	}
	secondAddress := Address{
		UserID:  456,
		Address: "Tokopedia Tower Lt 52",
		ZipCode: "12950",
	}
	generateShipment(&gtl, secondAddress)
}

func generateShipment(logistic Logistic, address Address) {
	logistic.CreateShipment(address)
}
