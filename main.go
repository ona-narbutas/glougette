package main

import (
	"fmt"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ona-narbutas/glougette/internal/inventory"
)

// "os"

// func homePageHandler(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("./static/home.html")
// 	if (err != nil) {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}
// 	t.ExecuteTemplate(w, "home.html", nil)
// }

// MAIN
func main() {
	inv := inventory.CreateInventory()
	inv.AddBrand("Free People")
	inv.AddBrand("Nic + Zoe")
	fmt.Println("inv", inv)
	fmt.Println(inventory.SaveInventoryData(*inv))

	// Create router
	router := gin.Default()
	router.GET("/")

	// Routes
	// http.HandleFunc("/", homePageHandler)
	// router.GET("/", homePageHandler)
	router.GET("/inventory", getInventory)

	// Start server on port 8080
	// log.Fatal(http.ListenAndServe(":8080", nil))
	router.Run("localhost:8080")
}

func getInventory(c *gin.Context) {
	inv, err := inventory.RetrieveInventory()
	if err != nil {
		fmt.Println("Retrieve Inventory Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "inventory not found: "})
		return
	}
	c.IndentedJSON(http.StatusOK, inv)
}