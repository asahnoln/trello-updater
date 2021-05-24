package trello

import (
	"flag"
	"os"
	"testing"
)

var baseUrl string
func init() {
    flag.StringVar(&baseUrl, "url", "", "API base URL")
}

func TestMain(m *testing.M) {
    flag.Parse()
    os.Exit(m.Run())
}

func TestGetLists(t *testing.T) {
	lists, err := GetLists(baseUrl)
	if err != nil {
		t.Errorf("got err %w, want no error", err)
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

func TestGetListsWithError(t *testing.T) {
}
