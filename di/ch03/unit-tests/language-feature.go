package unit_tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Pet struct {
	Name string
}

func NewPet(name string) *Pet {
	return &Pet{
		Name:name,
	}
}

func TestLanguageFeatures(t *testing.T)  {
	petFish := NewPet("Goldie")
	assert.IsType(t, &Pet{}, petFish)
}
