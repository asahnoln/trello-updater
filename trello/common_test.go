package trello

import (
	"flag"
	"os"
	"testing"
)

var baseUrl string
var boardId int

func init() {
	flag.StringVar(&baseUrl, "url", "", "API base URL")
	flag.IntVar(&boardId, "boardId", 0, "API board ID")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
