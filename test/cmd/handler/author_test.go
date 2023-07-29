package handler

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"simple-bookshelf/cmd/author"
	"simple-bookshelf/cmd/handler"
	"simple-bookshelf/test/cmd/mock"
	"testing"
)

func TestGetAuthorById(t *testing.T) {
	// Setup
	r := SetupRouter()
	// Setup Handler
	newService := author.NewService()
	authorHandler := handler.NewAuthorHandler(newService)

	r.GET("/authors/:id", authorHandler.GetAuthorById)

	// Test Case List
	testCases := []struct {
		description string
		input       string
		expect      author.Author
	}{
		{
			description: "Should return author data",
			input:       "1",
			expect:      mock.GenerateAuthor(""),
		},
	}

	for _, testData := range testCases {

		t.Run(testData.description, func(t *testing.T) {
			// Given

			// When
			req, _ := http.NewRequest("GET", fmt.Sprintf("/authors/%s", testData.input), nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Then
			var result author.Author
			err := json.Unmarshal(w.Body.Bytes(), &result)
			if err != nil {
				t.Errorf("Expect %+v, got error %s", testData.expect, err.Error())
			}

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, testData.expect, result)
		})

	}
}
