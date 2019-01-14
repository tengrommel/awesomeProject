package main

import (
	"fmt"
	"go.planetmeican.com/titan/log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	startString := "2016-01-01"
	endString := "2016-02-15"
	q := url.Values{}
	q.Set("cityId", "1")
	q.Set("start", startString)
	q.Set("end", endString)
	sb := strings.Builder{}
	sb.WriteString("https://6u7xxiivah.execute-api.cn-north-1.amazonaws.com.cn/prod/play/accident/result/computed/cityid")
	sb.WriteString("?")
	sb.WriteString(q.Encode())
	fmt.Println(sb.String())
	//req, err := http.NewRequest(http.MethodGet, "https://6u7xxiivah.execute-api.cn-north-1.amazonaws.com.cn/prod/internalapi/play/accident/result/computed/cityid?cityId=1&end=2016-02-26&start=2016-01-01", nil)
	req, err := http.NewRequest(http.MethodGet, "https://6u7xxiivah.execute-api.cn-north-1.amazonaws.com.cn/prod/internalapi/play/accident/result/computed/cityid?cityId=1&end=2016-02-15&start=2016-01-01", nil)
	if err != nil {
		log.Error(err)
		return
	}
	req.Header.Set("Authorization", "Bearer FWO01RyjG7LdUf2vC9oMwLQE1P9t6GhFJt3KwAS7NRzbz9qReH")
	req.Header.Set("x-api-key", "xFgl7kDs5u9hkEu2mOGPL2S1GpVvvBZi7dUs2bu4")
	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		return
	}
	fmt.Println(resp)
}
