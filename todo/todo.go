package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Text     string
	Priority int
	Position int
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
	for i, _ := range items {
		items[i].Position = i + 1
	}
	return items, nil
}

func (i *Item) SetPriority(pri int) { // This is an example of writing a method for a reciever
	// The receiver here is of type Item
	fmt.Println("Came in SetPriority with priority", pri)
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	switch i.Priority {
	case 1:
		return "(1)"
	case 3:
		return "(3)"
	default:
		return " "
	}
}

func (i *Item) GetPosition() string {
	return fmt.Sprintf("%d", i.Position)
}

// ByPri implements sort.Interface for []Item based on the priority and Position
type ByPri []Item

func (s ByPri) Len() int {
	return len(s)
}

// This is a swap by the Position  becuase labnguage is based on pointers
func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].Position < s[j].Position
	}
	return s[i].Priority < s[j].Priority
}
