package trello

import (
	"testing"
)

func TestGetCards(t *testing.T) {
	checkFlags(t)

	lists := listsForTests(t)
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
		case c.Id == "":
			t.Errorf("c.Id = \"\", want non-zero value")
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
