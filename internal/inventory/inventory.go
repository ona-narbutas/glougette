package inventory

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// TYPES
type Inventory struct {
	Brands *[]Brand 					`json:"brands"`
}

type Brand struct {
	Name string 							`json:"name"`
	Lines *[]Line 						`json:"lines"`
}

type Line struct {
	Name string 							`json:"name"`
	Items *[]Item							`json:"items"`
}

type Item struct {
	Name string								`json:"name"`
	Category string						`json:"category"`
	Cost float64							`json:"cost"`
	Description string				`json:"description"`
	Price float64							`json:"price"`
	Instances *[]ItemInstance	`json:"instances"`
}

type ItemInstance struct {
	SizeSML string						`json:"sizeSML"`
	SizeNumeric int						`json:"sizeNumeric"`
	Color string							`json:"color"`
	AtStore string						`json:"atStore"`
}

// METHODS

// Get quantities
func (i *Item) GetItemQuantity() int {
	return len(*i.Instances)
}

func (l *Line) GetLineQuantity() int {
	q := 0
	for _, i := range *l.Items {
		q += i.GetItemQuantity()
	}
	return q
}

func (b *Brand) GetBrandQuantity() int {
	q := 0
	for _, l := range *b.Lines {
		q += l.GetLineQuantity()
	}
	return q
}

// Add brands
func (i *Inventory) AddBrand(name string) error {
	for _, b := range *i.Brands {
		if b.Name == name {
			return errors.New("Brand already exists")
		}
	}
	lines := make([]Line, 0)
	*i.Brands = append(*i.Brands, Brand{name, &lines})
	fmt.Println(*i.Brands)
	return nil
}


// INVENTORY CREATION

// Create new inventory
func CreateInventory() *Inventory {
	brands := make([]Brand, 0)
	return &Inventory{&brands}
}

// Save current inventory state to json
func SaveInventoryData(inv Inventory) error {
	fmt.Println("formatting data", inv)
	data, err := json.Marshal(inv)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("writing data", string(data))
	return os.WriteFile("data/inventory.json", data, 0644)
}