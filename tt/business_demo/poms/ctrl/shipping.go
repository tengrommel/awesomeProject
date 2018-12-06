package ctrl

import (
	"awesomeProject/tt/business_demo/poms/model"
	"encoding/json"
	"net/http"
)

type ShippingController struct{}

func (sc *ShippingController) GetReceivers(w http.ResponseWriter, r *http.Request) {
	receivers := model.GetCurrencies()
	w.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(receivers)
	w.Write(data)
}

func (sc *ShippingController) GetVendors(w http.ResponseWriter, r *http.Request) {
	vendors := model.GetVendors()
	w.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(vendors)
	w.Write(data)
}
