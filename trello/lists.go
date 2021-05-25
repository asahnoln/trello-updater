package trello

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type List struct {
	Id, Name string
	Closed   bool
}

type ListSlice []*List

func GetLists(baseUrl string, boardId string) (ListSlice, error) {
	res, err := http.Get(baseUrl + "/boards/" + boardId + "/lists")
	if err != nil {
		return nil, fmt.Errorf("trello: %w", err)
	}
	defer res.Body.Close()

	listsJson, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("trello: %w", err)
	}

    lists, err := Lists(listsJson)
	return lists, err
}

func Lists(data []byte) (ListSlice, error) {
	var lists ListSlice
    err := json.Unmarshal(data, &lists)
    if err != nil {
        return lists, err
    }

	return lists, nil
}
