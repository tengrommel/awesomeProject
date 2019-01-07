package main

import (
	"fmt"
	"reflect"
)

type MonthlyExportByCityIdResultRow struct {
	CorpId                  string `header:"CorpId"`
	CorpName                string `header:"CorpName"`
	RestaurantId            string `header:"RestaurantId"`
	RestaurantName          string `header:"RestaurantName"`
	DeliveryFeeInCent       string `header:"DeliveryFeeInCent"`
	PayableInCent           string `header:"PayableInCent"`
	UserToMeicanPriceInCent string `header:"UserToMeicanPriceInCent"`
	DiscountRatio           string `header:"DiscountRatio"`
	Receivable              string `header:"Receivable"`
	UseSpecialAccount       string `header:"UseSpecialAccount"`
}

func CreateHeaderStringSlice(ex interface{}) ([]string, error) {
	res := make([]string, 0)
	exType := reflect.TypeOf(ex)
	exValue := reflect.ValueOf(ex)
	kd := exValue.Kind()
	if kd == reflect.Ptr {
		err := fmt.Errorf("can not be ptr")
		return res, err
	}
	if exType.Kind() != reflect.Struct {
		err := fmt.Errorf("must be struct")
		return res, err
	}
	structLen := exValue.NumField()
	for i := 0; i < structLen; i++ {
		item := exType.Field(i).Tag.Get("header")
		if item != "" {
			res = append(res, exType.Field(i).Tag.Get("header"))
		}
	}
	return res, nil
}

func main() {
	m := MonthlyExportByCityIdResultRow{
		CorpId: "1",
	}
	list, _ := CreateHeaderStringSlice(m)
	fmt.Println(list)
}
