package trello

import (
	"testing"
)

func TestGetCards(t *testing.T) {
	if baseUrl == "" {
		t.Fatalf("-url flag should be set to API url")
	}
	if boardId == 0 {
		t.Fatalf("-boardId flag should be set to board ID")
	}

	lists, err := GetLists(baseUrl, boardId)
	if err != nil {
		t.Errorf("GetLists() = got err %w, want no error", err)
	}
	cards, err := GetCards(baseUrl, lists[0].Id)
	if err != nil {
		t.Errorf("GetCards() = got err %w, want no error", err)
	}

	switch {
	case cards == nil:
		t.Fatalf("GetCards() = nil, want slice")
	case len(cards) == 0:
		t.Fatalf("GetCards() = 0, want more than 0")
	}

	for _, c := range cards {
		switch {
		case c.Id == 0:
			t.Errorf("c.Id = 0, want non-zero value")
		case c.Name == "":
			t.Errorf("c.Name = \"\", want non-zero value")
		case c.Desc == "":
			t.Errorf("c.Desc = \"\", want non-zero value")
		}
	}
}

// TODO: Errors should be handled properly
func TestGetCardsWithError(t *testing.T) {
}
