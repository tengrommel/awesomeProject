package parser

import (
	"awesomeProject/m/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "你")
	if len(result.Items) != 1 {
		t.Errorf("Item should contain 1"+
			"element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:   "你",
		Gender: "女",
	}
	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
