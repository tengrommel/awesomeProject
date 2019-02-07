package routes

import (
	"awesomeProject/notice/go_web/utils"
	"net/http"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}
