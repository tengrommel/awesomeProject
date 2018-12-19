package main

import (
	"fmt"
	"github.com/speps/go-hashids"
	"reflect"
)

var salt = "Df3FAnyzwoNUGQ6Rd63wBwURUBGesqs9YNMNHPLnVkLyk77qaV"

type Ex struct {
	NameId int64  `encrypt:"true"`
	NoId   string `encrypt:"true"`
	NotId  string
}

func EncryptID(ex interface{}) (err error) {
	var (
		exType    reflect.Type
		exValue   reflect.Value
		kd        reflect.Kind
		structLen int
		value     string
	)
	exType = reflect.TypeOf(ex)
	exValue = reflect.ValueOf(ex)
	kd = exValue.Kind()
	if kd != reflect.Ptr {
		return fmt.Errorf("must be struct ptr")
	}
	if exValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("must be struct ptr")
	}
	structLen = exValue.Elem().NumField()
	for i := 0; i < structLen; i++ {
		if exType.Elem().Field(i).Type.String() != "string" {
			continue
		}
		if exType.Elem().Field(i).Tag.Get("encrypt") == "true" {
			value, err = encryptID(exValue.Elem().Field(i).String())
			if err != nil {
				return
			}
			exValue.Elem().Field(i).SetString(value)
		}
	}
	return
}

func encryptID(str string) (encryptString string, err error) {
	data := hashids.NewData()
	data.Salt = salt
	data.Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"
	data.MinLength = 15
	var hashid *hashids.HashID
	hashid, err = hashids.NewWithData(data)
	encryptString, err = hashid.EncodeHex(str)
	return
}

func main() {
	example := Ex{
		NameId: 1,
		NoId:   "4",
	}
	EncryptID(&example)
	fmt.Println(example)
}
