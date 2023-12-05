package main

import "fmt"

type Product struct {
	ID    int64
	Name  string
	Price float64
	Stock int
}

func main() {
	product := Product{
		ID:    1,
		Name:  "Jacket",
		Price: 150000,
		Stock: 0,
	}

	isOutOfStock := product.IsOutOfStock()
	fmt.Println(isOutOfStock)
}

func (p *Product) IsOutOfStock() bool {
	return p.Stock == 0
}
