package mdl

import (
	"fmt"
	models "mdl/lib/models"
	utils "mdl/lib/utils"
	"net/http"
	"os"
	"strings"
)

var data models.Data

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	category, err := os.ReadFile("./data/category.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data.Category = strings.Split(string(category), ",")
	for i := range [6]int{} {
		book := models.Book{
			BookImg: "../assets/img_books" + fmt.Sprintf("%d", i),
			BookName: fmt.Sprintf("%d", i),
		}
		data.Books = append(data.Books, book)
	}
	utils.RenderPage("index", data, w)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderPage("about", nil, w)
}

func AuthorHandler(w http.ResponseWriter, r *http.Request) {

}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	
}

func CitationHandler(w http.ResponseWriter, r *http.Request) {
	
}

func ConditionHandler(w http.ResponseWriter, r *http.Request) {
	
}

