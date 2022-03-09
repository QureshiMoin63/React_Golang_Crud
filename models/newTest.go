// package main

// import (
// 	"bytes"
// 	"net/http"
// 	"testing"
// )

// func TestRegister(t *testing.T) {

// 	var jsonStr = []byte(`{"title":"deepak","email":"xyz@gmail.com","password":"123456"`)

// 	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// }
// func TestLogin(t *testing.T) {

// 	var jsonStr = []byte(`{"email":"xyz@gmail.com","password":"123456"`)

// 	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// }

package models

import (
	"encoding/json"
	"testing"
)

func TestAuthorValidations(t *testing.T) {

	TestValues := []string{
		`{"id":1, "name":"Anna Karenina","description":"greatest novels"}`,
	}

	for _, value := range TestValues {

		var testAuthor Author
		err := json.Unmarshal([]byte(value), &testAuthor)
		if err != nil {
			t.Errorf("Test Failed : %v", err.Error())
		}

		err = testAuthor.Validate()
		if err != nil {
			t.Errorf("Test Failed : %v", err.Error())
		}
	}
}

// func TestBookValidations(t *testing.T) {

// 	TestValues := []string{
// 		`{"id": 1, "name": "greatest novels","description":"Any fan of stories that involve juicy subjects like adultery, gambling, marriage plots, and, well,Russian feudalism, would instantly place Anna Karenina at the peak of their “greatest novels” list. And that’s exactly the ranking that publications like Time magazine have given the novel since it was published in its entirety in 1878."}}`,

// 		`{"id": 2, "name": "Harper Lee", "description": "Harper Lee, believed to be one of the most influential authors to have ever existed, famously published only a single novel (up until its controversial sequel was published in 2015 just before her death)."}`,
// 	}

// 	for _, value := range TestValues {

// 		var testBook models.Book

// 		if err := json.Unmarshal([]byte(value), &testBook); err != nil {
// 			t.Errorf("Failed to unmarshal %q To Json : %v", value, err.Error())
// 		}

// 		if err := testBook.Validate(); err != nil {
// 			t.Errorf("Test Failed : %v", err.Error())
// 		}
// 	}
// }

// func TestUserValidations(t *testing.T) {

// 	TestValues := []string{
// 		`{"username": "hello world", "email": "Moin qureshi", "password": "$2a$10$xYADUPksQV3kPzis5I0.ruBlgzmxaTx9uDJSHkxXfNB..."}`,
// 	}

// 	for i, value := range TestValues {

// 		var testUser models.User // A new User type variable

// 		if err := json.Unmarshal([]byte(value), &testUser); err != nil {
// 			t.Errorf("Failed to unmarshal %q To Json : %v", value, err.Error())

// 		}

// 		if err := testUser.Validate(); err != nil {
// 			t.Errorf("(%v) => Test Failed : %v", i, err.Error()) // index starts with 0 in a slice.
// 		}
// 	}
// }
