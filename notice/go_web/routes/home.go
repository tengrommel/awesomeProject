package routes

import (
	"awesomeProject/notice/go_web/models"
	"awesomeProject/notice/go_web/utils"
	"net/http"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	utils.ExecuteTemplate(w, "home.html", struct {
		Categories []models.Category
	}{
		Categories: categories,
	})
}
