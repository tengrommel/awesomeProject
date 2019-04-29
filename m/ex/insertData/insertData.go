package insertData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Data []struct {
		Type       string `json:"type"`
		Attributes struct {
			TeamOne      string `json:"teamOne"`
			TeamTwo      string `json:"teamOne"`
			TeamOneGoals int    `json:"teamOneGoals"`
			TeamTwoGoals int    `json:"teamTwoGoals"`
		} `json:"attributes"`
	} `json:"data"`
}

type Error struct {
	Status string `json:"status"`
	Source struct {
		Pointer string `json:"pointer"`
	} `json:"source"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type SuccessMessage struct {
	Meta struct {
		Success struct {
			Title   string `json:"title"`
			Message string `json:"message"`
			Status  string `json:"status"`
		} `json:"success"`
	} `json:"meta"`
}

func insertData(r *http.Request) ([]byte, error) {
	var result Result
	b, _ := ioutil.ReadAll(r.Body) // reads the body of the request to /insert
	err := json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println(err)
	}
	// checks if the request body is correct
	isSuccess, message := isJSONCorrect(result)

	if isSuccess {
		return getSuccessJsonBody("Game result was successfully inserted"), nil
	} else {
		return getErrorFromResult("/insert", "Invalid Attribute", message), nil
	}
}
