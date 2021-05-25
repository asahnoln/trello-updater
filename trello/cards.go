package trello

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Card struct {
	Id, Name, Desc string
	Closed         bool
}

func GetCards(baseUrl string, listId string) ([]*Card, error) {
	res, err := http.Get(baseUrl + "/lists/" + listId + "/cards")
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
