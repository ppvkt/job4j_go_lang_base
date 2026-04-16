package tracker

import (
	"strings"
)

type ItemCrud interface {
	AddItem(item Item) (Item, error)
	GetItems() []Item
	UpdateItem(uuid string, newName string) error
	DeleteItem(uuid string) error
	FindItem(namePart string) []Item
}

type Tracker struct {
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) (Item, error) {
	_, ok := t.indexOf(item.ID)
	if ok {
		return item, ErrHasAlreadyExist
	}

	t.Items = append(t.Items, item)

	return item, nil

}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.Items))
	copy(res, t.Items)
	return res
}

func (t *Tracker) UpdateItem(uuid string, newName string) error {
	index, ok := t.indexOf(uuid)

	if !ok {
		return ErrNotFound
	}

	t.Items[index].Name = newName

	return nil
}

func (t *Tracker) DeleteItem(uuid string) error {
	index, ok := t.indexOf(uuid)

	if !ok {
		return ErrNotFound
	}

	copy(t.Items[index:], t.Items[index+1:])
	t.Items = t.Items[:len(t.Items)-1]

	return nil
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

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.Items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
}
