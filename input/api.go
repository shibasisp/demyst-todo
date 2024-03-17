package input

import (
	"demyst-todo/log"
	"demyst-todo/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type API struct {
	URL string
}

func (a *API) Fetch(limit int) ([]types.TODO, error) {
	numRequests := limit
	concurrency := 10

	errorsCh := make(chan error, concurrency)
	respChan := make(chan *types.TODO, concurrency)
	defer close(respChan)
	defer close(errorsCh)
	todos := make([]types.TODO, 0, limit)
	var wg sync.WaitGroup

	for i := 1; i <= numRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			a.fetchTodo(id, errorsCh, respChan)
		}(i)
	}

	go func() {
		wg.Wait()
	}()

	for i := 0; i < numRequests; i++ {
		select {
		case todo, ok := <-respChan:
			if !ok {
				log.Logger.Errorf("Closed resp channel")
				continue
			}
			if todo != nil {
				log.Logger.Infof("Fetched TODO: ID: %d, Title: %s, Completed: %v\n", todo.ID, todo.Title, todo.Completed)
				todos = append(todos, *todo)
			}
		case err, ok := <-errorsCh:
			if !ok {
				log.Logger.Errorf("Closed err channel")
				continue
			}
			if err != nil {
				log.Logger.Errorf("Error fetching todos from API, error: %v", err)
				return todos, err
			}
		}
	}

	return todos, nil
}

func (a *API) fetchTodo(id int, errChan chan<- error, respChan chan<- *types.TODO) {
	url := fmt.Sprintf("%s/%d", a.URL, id)
	log.Logger.Info("Calling external API with id :", id)
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
	if resp.StatusCode >= 400 {
		errChan <- fmt.Errorf("error fetching TODO from API, status code: %d", resp.StatusCode)
		return
	}
	var todo types.TODO
	if err := json.Unmarshal(body, &todo); err != nil {
		errChan <- err
		return
	}

	respChan <- &todo
}
