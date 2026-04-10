package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// produsct
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/edit", productcontroller.Edit)
	http.HandleFunc("/product/delete", productcontroller.Delete)

	log.Println("server running on port 8383")
	http.ListenAndServe(":8383", nil)
}
