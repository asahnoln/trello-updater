package trello

import (
	"testing"
)

func TestGetLists(t *testing.T) {
	if baseUrl == "" {
		t.Fatalf("-url flag should be set to API url")
	}

	lists, err := GetLists(baseUrl)
	if err != nil {
		t.Errorf("GetLists() = got err %w, want no error", err)
	}

	switch {
	case lists == nil:
		t.Errorf("GetLists() = nil, want slice")
	case len(lists) == 0:
		t.Errorf("GetLists() = 0, want more than 0")
	}

	for _, l := range lists {
		switch {
		case l.Id == 0:
			t.Errorf("l.Id = 0, want non-zero value")
		case l.Name == "":
			t.Errorf("l.Name = \"\", want non-zero value")
		}
	}
}

// TODO: Errors should be handled properly
func TestGetListsWithError(t *testing.T) {
}
