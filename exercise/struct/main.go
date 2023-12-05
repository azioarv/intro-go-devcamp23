package main

import (
	"fmt"
)

type Review struct {
	UserID     int64
	Content    string
	ImagesPath []string
	ProductID  int64
	Rating     int
}

func main() {
	review := Review{
		UserID:     123,
		Content:    "Barangnya bagus",
		ImagesPath: []string{"images_1.jpg"},
		ProductID:  4,
		Rating:     5,
	}

	fmt.Printf("%#v", review)
}
