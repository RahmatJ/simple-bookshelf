package handler

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"simple-bookshelf/cmd/book"
	"simple-bookshelf/cmd/handler"
	"simple-bookshelf/test/cmd/mock"
	"testing"
)

func TestGetBook(t *testing.T) {
	//	Setup
	r := SetupRouter()

	//	Setup Handler
	bookHandler := handler.BookHandler{}

	r.GET("/books", bookHandler.GetBook)

	//	Test Case List
	testCases := []struct {
		description string
		expect      book.Book
	}{
		{
			description: "Should return Book data",
			expect:      mock.GenerateBook(),
		},
	}

	for _, testData := range testCases {
		t.Run(testData.description, func(t *testing.T) {
			//	Given
			//	When
			req, _ := http.NewRequest("GET", fmt.Sprintf("/books"), nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			//	Then
			var result book.Book
			err := json.Unmarshal(w.Body.Bytes(), &result)
			if err != nil {
				t.Errorf("Expect %+v, got error %s", testData.expect, err.Error())
			}

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, testData.expect, result)
		})
	}
}
