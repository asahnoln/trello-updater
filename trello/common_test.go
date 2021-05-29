package trello

import (
	"flag"
	"os"
	"testing"
)

var baseUrl string
var boardId string

func init() {
	flag.StringVar(&baseUrl, "url", "", "API base URL")
	flag.StringVar(&boardId, "boardId", "", "API board ID")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func checkFlags(t *testing.T) {
    switch {
    case baseUrl == "":
		t.Fatalf("-url flag should be set to API url")
    case boardId == "":
		t.Fatalf("-boardId flag should be set to board ID")
    }
}
