package productcontroller

import (
	"go-web-native/models/productmodel"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {

}
func Edit(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
