package input_test

import (
	"demyst-todo/input"
	"demyst-todo/types"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func compareMyStruct(a, b types.TODO) bool {
	return a.ID == b.ID && a.UserID == b.UserID && a.Title == b.Title && a.Completed == b.Completed
}

func TestFetchWithSuccess(t *testing.T) {
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
	api := input.API{URL: "http://example.com/todos"}

	result, err := api.Fetch(2)
	callCount := httpmock.GetTotalCallCount()
	assert.Equal(t, 2, callCount)

	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, result, compareMyStruct)
}

func TestFetchWithBadRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		`http://example.com/todos/1`,
		httpmock.NewStringResponder(400, "{}"))

	api := input.API{URL: "http://example.com/todos"}

	result, err := api.Fetch(1)
	callCount := httpmock.GetTotalCallCount()
	assert.Equal(t, 1, callCount)

	assert.NotNil(t, err)
	assert.Equal(t, []types.TODO{}, result)
}

func TestFetchWithError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		`http://example.com/todos/1`,
		httpmock.NewErrorResponder(fmt.Errorf("err")))

	api := input.API{URL: "http://example.com/todos"}

	result, err := api.Fetch(1)
	callCount := httpmock.GetTotalCallCount()
	assert.Equal(t, 1, callCount)

	assert.NotNil(t, err)
	assert.Equal(t, []types.TODO{}, result)
}

type MockReader struct{}

func (m MockReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mock error")
}
func (m MockReader) Close() error {
	return errors.New("mock error")
}

func TestFetchWithInvalidResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	reader := MockReader{}
	httpmock.RegisterResponder("GET", "http://example.com/todos/1", func(req *http.Request) (*http.Response, error) {
		response := httpmock.NewStringResponse(http.StatusOK, "")
		response.Body = reader
		return response, nil
	})

	api := input.API{URL: "http://example.com/todos"}

	result, err := api.Fetch(1)
	callCount := httpmock.GetTotalCallCount()
	assert.Equal(t, 1, callCount)

	assert.NotNil(t, err)
	assert.Equal(t, []types.TODO{}, result)
}

func TestFetchWithErrorUnmarshal(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		`http://example.com/todos/1`,
		httpmock.NewBytesResponder(200, []byte{123}))

	api := input.API{URL: "http://example.com/todos"}

	result, err := api.Fetch(1)
	callCount := httpmock.GetTotalCallCount()
	assert.Equal(t, 1, callCount)

	assert.NotNil(t, err)
	assert.Equal(t, []types.TODO{}, result)
}
