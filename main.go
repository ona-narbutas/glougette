package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ona-narbutas/glougette/internal/inventory"
)

// "os"



func homePageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/home.html")
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	t.ExecuteTemplate(w, "home.html", nil)
}

// MAIN
func main() {
	inv := inventory.CreateInventory()
	inv.AddBrand("Free People")
	inv.AddBrand("Nic + Zoe")
	fmt.Println("inv", inv)
	fmt.Println(inventory.SaveInventoryData(*inv))
	// Routes
	http.HandleFunc("/", homePageHandler)

	// Start server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}