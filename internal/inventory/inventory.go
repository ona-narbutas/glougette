package inventory

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// TYPES
type Inventory struct {
	Brands map[string]Brand 	`json:"brands"`
}

type Brand struct {
	Lines map[string]Line 		`json:"lines"`
}

type Line struct {
	Items map[string]Item			`json:"items"`
}

type Item struct {
	Name string								`json:"name"`
	Category string						`json:"category"`
	Cost float64							`json:"cost"`
	Description string				`json:"description"`
	Price float64							`json:"price"`
	Instances *[]ItemInstance	`json:"instances"`
	HasNumericSize bool				`json:"hasNumericSize"`
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
	for _, i := range l.Items {
		q += i.GetItemQuantity()
	}
	return q
}

func (b *Brand) GetBrandQuantity() int {
	q := 0
	for _, l := range b.Lines {
		q += l.GetLineQuantity()
	}
	return q
}

// Add brand
func (i *Inventory) AddBrand(name string) error {
	_, ok := i.Brands[name]
	if ok == true {
		return errors.New("Brand already exists")
	}

	brand := make(map[string]Line)
	i.Brands[name] = Brand{brand}
	return nil
}

// Add Line
func (b *Brand) AddLine(name string) error {
	_, ok := b.Lines[name]
	if ok == true {
		return errors.New("Line already exists")
	}

	line := Line{make(map[string]Item)}
	b.Lines[name] = line
	return nil
}

// Add item type
func (inv *Inventory) AddItemToInventory(brand string, line string, item Item) error {
	inv.AddBrand(brand)
	inv.Brands[brand].addItemToBrand(line, item)
	return nil
}

func (b Brand) addItemToBrand(line string, item Item) error {
	b.AddLine(line)
	b.Lines[line].addItemToLine(item)
	return nil
}

func (l Line) addItemToLine(item Item) error {
	_, ok := l.Items[item.Name]
	if ok == true {
		return errors.New("Item already exists")
	}

	l.Items[item.Name] = item
	return nil
}


// INVENTORY CREATION

// Create new inventory
func CreateInventory() *Inventory {
	brands := make(map[string]Brand)
	return &Inventory{brands}
}

// Save current inventory state to json
func SaveInventoryData(inv Inventory) error {
	data, err := json.Marshal(inv)
	if err != nil {
		return err
	}
	return os.WriteFile("data/inventory.json", data, 0644)
}

func RetrieveInventory() (Inventory, error) {
	invData, err := os.ReadFile("data/inventory.json")
	fmt.Println("read data", invData)
	if err != nil {
		return Inventory{}, err
	}

	var inv *Inventory = &Inventory{}
	if err := json.Unmarshal(invData, inv); err != nil {
		return Inventory{}, err
	}
	return *inv, nil
}