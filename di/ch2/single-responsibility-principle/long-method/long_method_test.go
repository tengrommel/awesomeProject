package long_method

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestLoadUserHandler(t *testing.T)  {
	// build request
	req := &http.Request{
		Form: url.Values{},
	}
	req.Form.Add("UserID", "1234")
	// call function under test
	resp := httptest.NewRecorder()
	loadUserHandlerLong(resp, req)
	// validate result
	assert.Equal(t, http.StatusOK, resp.Code)
	expectedBody := `{"ID":1,"Name":"Bob","Phone":"0123456789"}` + "\n"
	assert.Equal(t, expectedBody, resp.Body.String())
}

func TestMain(m *testing.M)  {
	var mock sqlmock.Sqlmock
	DB, mock, _ = sqlmock.New()
	mock.ExpectQuery(".*").WillReturnRows(
		sqlmock.NewRows([]string{"ID", "Name", "Phone"}).AddRow(
			1, "Bob", "0123456789"))
	os.Exit(m.Run())
}