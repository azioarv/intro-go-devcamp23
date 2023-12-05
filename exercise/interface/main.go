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
	TrackShipment(userID int64) Response
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

func (g *GoSend) TrackShipment(userID int64) Response {
	// api call process
	fmt.Println("Track Shipment for Go-Send For UserID ", userID)
	return Response{}
}
func (gtl *GTL) TrackShipment(userID int64) Response {
	// api call process
	fmt.Println("Track Shipment for GTL For UserID ", userID)
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
	trackShipment(&gosend, firstAddress.UserID)

	gtl := GTL{
		URL: "https://api.gtl.id",
	}
	secondAddress := Address{
		UserID:  456,
		Address: "Tokopedia Tower Lt 52",
		ZipCode: "12950",
	}
	generateShipment(&gtl, secondAddress)
	trackShipment(&gtl, secondAddress.UserID)
}

func generateShipment(logistic Logistic, address Address) {
	logistic.CreateShipment(address)
}

func trackShipment(logistic Logistic, userID int64) {
	logistic.TrackShipment(userID)
}
