package input_test

import (
	"demyst-todo/input"
	"demyst-todo/types"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func compareMyStruct(a, b types.TODO) bool {
	return a.ID == b.ID && a.UserID == b.UserID && a.Title == b.Title && a.Completed == b.Completed
}

func TestFetch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		`=~^http://example.com/todos/1`,
		httpmock.NewStringResponder(200, `{"id": 1, "userId": 1, "title": "Test TODO", "completed": false}`))

	httpmock.RegisterResponder(
		"GET",
		`=~^http://example.com/todos/2`,
		httpmock.NewStringResponder(200, `{"id": 2, "userId": 1, "title": "Test TODO 2", "completed": true}`))

	expected := []types.TODO{
		{ID: 1, UserID: 1, Title: "Test TODO", Completed: false},
		{ID: 2, UserID: 1, Title: "Test TODO 2", Completed: true},
	}
	api := input.API{URL: "http://example.com/todos/"}

	result, err := api.Fetch(2)
	callCount := httpmock.GetTotalCallCount()
	assert.Equal(t, 2, callCount)

	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result, compareMyStruct)
}
