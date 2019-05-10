package defining_depenency_injection

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSaver struct {
	mock.Mock
}

func (m *mockSaver) Save(data []byte) error {
	outputs := m.Mock.Called(data)
	return outputs.Error(0)
}

func TestSavePerson_happyPath(t *testing.T) {
	// input
	in := &Person{
		Name: "Sophia",
		Phone: "0123456789",
	}
	// mock the NFS
	mockNFS := &mockSaver{}
	mockNFS.On("Save", mock.Anything).Return(nil).Once()
	// Call Save
	resultErr := SavePerson(in, mockNFS)
	// validate result
	assert.NoError(t, resultErr)
	assert.True(t, mockNFS.AssertExpectations(t))
}
