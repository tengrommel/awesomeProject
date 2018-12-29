package prototype

import (
	"errors"
	"fmt"
)

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}

var (
	whitePrototype = &Shirt{
		Price: 15.00,
		SKU:   "empty",
		Color: White,
	}
	blackPrototype = &Shirt{
		Price: 16.00,
		SKU:   "empty",
		Color: Black,
	}
	bluePrototype = &Shirt{
		Price: 17.00,
		SKU:   "empty",
		Color: Blue,
	}
)

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}
