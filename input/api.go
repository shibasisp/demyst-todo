package input

import (
	"demyst-todo/log"
	"demyst-todo/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type API struct {
	URL string
}

func (a *API) Fetch(limit int) ([]types.TODO, error) {
	numRequests := limit
	concurrency := 10

	errorsCh := make(chan error, concurrency)
	respChan := make(chan *types.TODO, concurrency)
	defer close(errorsCh)
	defer close(respChan)
	todos := make([]types.TODO, 0, limit)

	//TODO: Implement Rate Limiter
	for i := 1; i <= numRequests; i++ {
		go a.fetchTodo(i, errorsCh, respChan)
	}

	for i := 0; i < numRequests; i++ {
		if err := <-errorsCh; err != nil {
			log.Logger.Errorf("Error fetching todos from API, error: %v", err)
		}
		if todo := <-respChan; todo != nil {
			log.Logger.Infof("Fetched TODO: ID: %d, Title: %s, Completed: %v\n", todo.ID, todo.Title, todo.Completed)
			todos = append(todos, *todo)
		}
	}

	return todos, nil
}

func (a *API) fetchTodo(id int, errChan chan<- error, respChan chan<- *types.TODO) {
	url := fmt.Sprintf("%s%d", a.URL, id)
	resp, err := http.Get(url)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}
	var todo types.TODO
	if err := json.Unmarshal(body, &todo); err != nil {
		errChan <- err
		return
	}
	respChan <- &todo
	errChan <- nil
}
