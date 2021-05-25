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

func checkFlags(t *testing.T) bool {
    switch {
    case baseUrl == "":
		t.Logf("-url flag is not set")
        return false
    case boardId == "":
		t.Logf("-boardId flag is not set")
        return false
    }

    return true
}
