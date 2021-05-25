package trello

import (
	"testing"
)

func listsForTests(t *testing.T) []*List {
	lists, err := GetLists(baseUrl, boardId)
	if err != nil {
		t.Errorf("GetLists() = got err %w, want no error", err)
	}

	switch {
	case lists == nil:
		t.Fatalf("GetLists() = nil, want slice")
	case len(lists) == 0:
		t.Fatalf("GetLists() = 0, want more than 0")
	}

	return lists
}

func TestGetLists(t *testing.T) {
	if !checkFlags(t) {
        t.Skipf("won't test http api without flags.")
    }

	lists := listsForTests(t)

	for _, l := range lists {
		switch {
		case l.Id == "":
			t.Errorf("l.Id = \"\", want non-zero value")
		case l.Name == "":
			t.Errorf("l.Name = \"\", want non-zero value")
		}
	}
}

func TestLists(t *testing.T) {
}

// TODO: Errors should be handled properly
func TestGetListsWithError(t *testing.T) {
}
