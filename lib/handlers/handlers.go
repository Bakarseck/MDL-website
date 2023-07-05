package mdl

import (
	models "mdl/lib/models"
	utils "mdl/lib/utils"
	"net/http"
	"os"
	"strings"
)

var data models.Data

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	category, err := os.ReadFile("category.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	data.Category = strings.Split(string(category), ",")
	utils.RenderPage("index", data, w)
}

