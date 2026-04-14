package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query("SELECT p.id, p.name, c.name, p.stock, p.description, p.created_at, p.updated_at FROM products p LEFT JOIN categories c ON p.category_id = c.id")
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
			&product.Category.Name,
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
func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`
		INSERT INTO products (name, category_id, stock, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)
	`, product.Name, product.Category.Id, product.Stock, product.Description, product.CreatedAt, product.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}
