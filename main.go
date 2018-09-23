package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/url"
)

func main() {
	urlJoin := url.URL{}
	urlJoin.Scheme = "https"
	urlJoin.Host = "meican.com"
	urlJoin.Path = "/invite/join"
	q := urlJoin.Query()
	q.Set("customerId", "fff")
	q.Set("uuid", "ook")
	urlJoin.RawQuery = q.Encode()
	fmt.Println(urlJoin.String())

	u1, err := uuid.NewV4()
	if err != nil {
		return
	}
	fmt.Println(u1)
}
