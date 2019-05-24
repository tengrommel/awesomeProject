package defining_depenency_injection

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tengrommel/awesomeProject/di/ch1/defining-depenency-injection/interface"
	"testing"
)

func TestSavePerson_nfsAlwaysFails(t *testing.T)  {
	// input
	in := &defining_depenency_injection.Person{
		Name: "Sophia",
		Phone: "0123456789",
	}
	// mock the NFS
	mockNFS := &mockSaver{}
	mockNFS.On("Save", mock.Anything).Return(errors.New("save failed")).Once()
	// Call Save
	resultErr := defining_depenency_injection.SavePerson(in, mockNFS)
	// validate result
	assert.Error(t, resultErr)
	assert.True(t, mockNFS.AssertExpectations(t))
}