package optimizing_for_humans

import (
	"io/ioutil"
	"net/http"
)

// Apply just enough abstraction
// Excessive abstraction leads to an excessive mental burden and excessive typing.
type myGetter interface {
	Get(url string) (*http.Response, error)
}

func TooAbstract(getter myGetter, url string) ([]byte, error) {
	resp, err := getter.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}