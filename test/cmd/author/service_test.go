package author

import (
	"github.com/go-playground/assert/v2"
	"simple-bookshelf/cmd/author"
	"simple-bookshelf/test/cmd/mock"
	"testing"
)

func TestGetAuthors(t *testing.T) {
	//	Test case list
	testCases := []struct {
		description string
		input       any
		expect      []author.Author
	}{
		{
			description: "Should get all authors",
			input:       nil,
			expect:      []author.Author{mock.GenerateAuthor("")},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.description, func(T *testing.T) {
			//	Given
			//	When
			authorService := author.NewService()
			result, err := authorService.GetAuthors()
			//	Then
			lengthResult := len(result)

			assert.Equal(t, lengthResult, 1)
			assert.Equal(t, err, nil)
		})

	}
}
