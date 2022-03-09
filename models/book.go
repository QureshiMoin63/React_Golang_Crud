package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name" validate :"required,min=2,max=30" example:"new1"`
	Description string `json:"description" validate:"omitempty,max=500" example:"new1 description"`
}

type _Book struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateBookRequest struct {
	Name        string `json:"book_name"`
	Description string `json:"descript"`
}

func (book *Book) Validate() error {
	err := validate.Struct(book)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}

// @Title Create Book
// @Description Creates Books & Returns a Book based on the request
// @Param request body _Book true "Create Book Request"
// @Router /books/create [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// database.Connection()
	// DB= database.GetDB()
	// DB.AutoMigrate(&Book{})
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	DB.Create(&book)
	json.NewEncoder(w).Encode(book)
}

// @Title Get All Book
// @Description Get All Books based on the request
// @Router /books [get]
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	var books []Book
	DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

// @Title Get _Book by ID
// @Description Get Books by ID based on the request
// @Param id path int true "Get Books by ID Request"
// @Router /books/{id} [get]
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Please Enter An Integer")
	} else {
		DB.First(&book, params["id"])
		json.NewEncoder(w).Encode(book)
	}

}

// @Title Update _Book By ID
// @Description Update Books based on the request
// @Param request body _Book true "Update Book Request"
// @Param id path int true "Update Book Request"
// @Router /books/{id} [patch]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Print Enter An Integer")
	} else {
		DB.First(&book, params["id"])
		json.NewDecoder(r.Body).Decode(&book)
		DB.Save(&book)
		json.NewEncoder(w).Encode(book)
	}

}

// @Title Delete Book By ID
// @Description Delete Books based on the request
// @Param id path int true "Delete Book Request"
// @Router /books/{id} [delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Please Enter An Integer")
	} else {
		DB.Delete(&book, params["id"])
		json.NewEncoder(w).Encode("This Author %v has been successfully deleted")
	}

}

// for index, item := range author {
// 	if item.id == params["id"] {
// 		DB.Delete(&author, params["id"])
// 		json.NewEncoder(w).Encode("The Author Has been successfully deleted")
// 	} else {
// 		fmt.Printf("Please Enter Correct params")
// 	}
// }
