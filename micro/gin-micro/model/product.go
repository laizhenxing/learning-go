package model

import "strconv"

type Product struct {
	ProdID   int    `json:"pid"`
	ProdName string `json:"pname"`
}

func GenerateProduct(n int) (prods []*Product) {
	for i := 0; i < n; i++ {
		prod := &Product{
			ProdID:   i,
			ProdName: "product " + strconv.Itoa(i+1),
		}

		prods = append(prods, prod)
	}

	return
}
