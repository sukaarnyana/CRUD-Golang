package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Id,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	println("jumlah", len(products))

	return products
}
