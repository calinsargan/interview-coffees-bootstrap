package repository

import (
	"errors"

	"github.com/calinsargan/interview-coffees-bootstrap/internal/coffee"
)

type MemStorageCoffee []coffee.Coffee

var memStorage = MemStorageCoffee{}

func (self MemStorageCoffee) GetUniqueID() int {
	if len(self) == 0 {
		return 1
	}

	return self[len(self)-1].ID + 1
}

type Mem struct {
}

func NewMem() Mem {
	return Mem{}
}

func (self Mem) Create(coffee coffee.Coffee) (coffee.Coffee, error) {
	coffee.ID = memStorage.GetUniqueID()
	// todo: #5 add item to slice

	return coffee, nil
}

func (self Mem) Get(coffeeID int) (coffee.Coffee, error) {
	// todo: #5 find item in slice by ID

	return coffee.Coffee{}, errors.New("not found")
}

func (self Mem) Update(coffee coffee.Coffee) error {
	// todo: #5 overwrite item in slice by ID

	return nil
}

func (self Mem) Delete(coffeeID int) error {
	// todo: #5 remove item from slice by ID

	return nil
}
