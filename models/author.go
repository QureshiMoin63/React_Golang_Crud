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

var validate = validator.New()
var DB *gorm.DB
var err error

type Author struct {
	gorm.Model
	Name        string `json:"name" validate:"required,min=2,max=30" example:"MOIN"`
	Description string `json:"description"`
}

func (author *Author) Validate() error {
	err := validate.Struct(author)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	// database.Connection()
	// DB = database.GetDB()
	// DB.AutoMigrate(&Author{})
	w.Header().Set("Content-Type", "application/json")
	var author Author
	json.NewDecoder(r.Body).Decode(&author)
	var check_author Author
	DB.Table("authors").Where("name = ?", author.Name).Scan(&check_author)
	if err != nil {
		fmt.Printf("Failed to Create an Author :%v", err)
	}

	if check_author.Name == "" {
		DB.Create(&author)
		fmt.Fprintln(w, "new user created")
		json.NewEncoder(w).Encode(author)
	} else {
		fmt.Fprintln(w, "user already exist")
	}

}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applciation/json")
	var authors []Author
	DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	fmt.Println(params)
	var author Author
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintln(w, "Please Enter Integer")
	} else {
		DB.First(&author, params["id"])
		json.NewEncoder(w).Encode(author)
	}
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	var author Author
	params := mux.Vars(r)
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintln(w, "Please Enter Integer")
	} else {
		DB.First(&author, params["id"])
		json.NewDecoder(r.Body).Decode(&author)
		DB.Save(&author)
		json.NewEncoder(w).Encode(author)
	}

}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	var author Author
	_, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintln(w, "Please Enter Integer")
	} else {
		DB.Delete(&author, params["id"])
		fmt.Fprintf(w, "The Author Has been successfully deleted")

	}
}
