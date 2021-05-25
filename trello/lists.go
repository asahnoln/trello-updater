package trello

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type List struct {
	Id, Name string
	Closed bool
}

func GetLists(baseUrl string, boardId string) ([]*List, error) {
	res, err := http.Get(baseUrl + "/boards/" + boardId + "/lists")
	if err != nil {
		return nil, fmt.Errorf("trello: %w", err)
	}
	defer res.Body.Close()

	listsJson, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("trello: %w", err)
	}

	var lists []*List
	err = json.Unmarshal(listsJson, &lists)

	return lists, nil
}
