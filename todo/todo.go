package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, error1 := json.Marshal(items)
	err := os.WriteFile(filename, b, 0644)
	if error1 != nil {
		return error1
	}
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil { // Special Syntax here for the if block
		return []Item{}, err
	}
	return items, nil
}
