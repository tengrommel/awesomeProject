package ctrl

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	home          *HomeController          = new(HomeController)
	shipping      *ShippingController      = new(ShippingController)
	currency      *CurrencyController      = new(CurrencyController)
	purchaseOrder *PurchaseOrderController = new(PurchaseOrderController)
)

func Setup() {
	createResourceServer()
	r := mux.NewRouter()
	r.HandleFunc("/api/receivers", shipping.GetReceivers)
	r.HandleFunc("/api/vendors", shipping.GetVendors)
	r.HandleFunc("/api/currencies", currency.GetCurrencies)
	r.HandleFunc("/api/purchaseOrders/{poNumber:\\d+}",
		purchaseOrder.GetOrder).Methods("GET")
	http.Handle("/", r)
}

func createResourceServer() {
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("res"))))
}
