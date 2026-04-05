package tracker

import "strings"

type ItemCrud interface {
	AddItem(item Item)
	GetItems() []Item
	UpdateItem(uuid string, newName string) bool
	DeleteItem(uuid string) bool
	FindItem(namePart string) []Item
}

type Tracker struct {
	Items []Item
}

type Item struct {
	ID   string
	Name string
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.Items = append(t.Items, item)
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) UpdateItem(uuid string, newName string) bool {
	for index, item := range t.Items {
		if uuid == item.ID {
			t.Items[index].Name = newName
			return true
		}
	}
	return false
}

func (t *Tracker) DeleteItem(uuid string) bool {
	for index, item := range t.Items {
		if uuid == item.ID {
			copy(t.Items[index:], t.Items[index+1:])
			t.Items = t.Items[:len(t.Items)-1]
			return true
		}
	}
	return false
}

func (t *Tracker) FindItem(namePart string) []Item {
	var result []Item
	namePart = strings.ToLower(namePart)

	for _, item := range t.Items {
		if strings.Contains(strings.ToLower(item.Name), namePart) {
			result = append(result, item)
		}
	}

	return result
}
