package trello

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Card struct {
	Id         int `json:",string"`
	Name, Desc string
	Closed     bool
}

func GetCards(baseUrl string, listId int) ([]*Card, error) {
	res, err := http.Get(baseUrl + "/lists/" + strconv.Itoa(listId) + "/cards")
	if err != nil {
		return nil, fmt.Errorf("trello: %w", err)
	}
	defer res.Body.Close()

	cardsJson, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("trello: %w", err)
	}

	var cards []*Card
	err = json.Unmarshal(cardsJson, &cards)

	return cards, nil
}
