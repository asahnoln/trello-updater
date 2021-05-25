package trello

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type List struct {
	Id     int `json:",string"`
	Name   string
	Closed bool
}

func GetLists(baseUrl string, boardId int) ([]*List, error) {
	res, err := http.Get(baseUrl + "/boards/" + strconv.Itoa(boardId) + "/lists")
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
