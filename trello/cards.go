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

type CardSlice []*Card

func GetCards(baseUrl string, listId string) (CardSlice, error) {
	res, err := http.Get(baseUrl + "/lists/" + listId + "/cards")
	if err != nil {
		return nil, fmt.Errorf("trello: %w", err)
	}
	defer res.Body.Close()

	cardsJson, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("trello: %w", err)
	}

	var cards CardSlice
	err = json.Unmarshal(cardsJson, &cards)

	return cards, nil
}
