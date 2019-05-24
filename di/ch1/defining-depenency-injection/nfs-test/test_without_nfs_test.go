package defining_depenency_injection

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tengrommel/awesomeProject/di/ch1/defining-depenency-injection/interface"
	"testing"
)

func TestSavePerson_happyPath(t *testing.T)  {
	// input
	in := &defining_depenency_injection.Person{
		Name: "Sophia",
		Phone: "0123456789",
	}
	mockNFS := &mockSaver{}
	mockNFS.On("Save", mock.Anything).Return(nil).Once()
	// Call Save
	resultErr := defining_depenency_injection.SavePerson(in, mockNFS)
	// validate result
	assert.NoError(t, resultErr)
	assert.True(t, mockNFS.AssertExpectations(t))
}

// mock implementation of Saver
type mockSaver struct {
	mock.Mock
}

// Save implement Saver
func (m *mockSaver)Save(data []byte) error {
	outputs := m.Mock.Called(data)
	return outputs.Error(0)
}