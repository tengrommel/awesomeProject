package defining_depenency_injection

import (
	"errors"
	"fmt"
)

// LoadPerson will load the requested person by ID.
// Errors include: invalid ID, missing person and failure to load or decode.
func LoadPerson(ID int, decodePerson func(data []byte) *Person) (*Person, error) {
	// validate the input
	if ID <= 0 {
		return nil, fmt.Errorf("invalid ID '%d' supplied", ID)
	}
	// load from storage
	bytes, err := loadPerson(ID)
	if err != nil {
		return nil, err
	}
	return decodePerson(bytes), nil
}

func loadPerson(ID int) ([]byte, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

