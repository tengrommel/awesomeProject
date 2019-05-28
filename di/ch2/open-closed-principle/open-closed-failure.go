package open_closed_principle

import (
	"io"
	"net/http"
)

// The first hint that something is amiss is the switch statement.
// It is not hard to imagine a situation where requirements change,
// and where we might need to add or even remove an output format.

func BuildOutputOCPFail(response http.ResponseWriter, format string, person Person)  {
	var err error
	switch format {
	case "csv":
		err = outputCSV(response, person)
	case "json":
		err = outputJSON(response, person)
	}
	if err != nil {
		// output a server error and quit
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}

// output the person as CSV and return error when failing to do so
func outputCSV(writer io.Writer, person Person) error {
	return nil
}

// output the person as JSON and return error when failing to do so
func outputJSON(writer io.Writer, person Person) error {
	return nil
}

// A data transfer object that represents a person
type Person struct {
	Name string
	Email string
}