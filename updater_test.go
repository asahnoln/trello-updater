package main

import (
	"flag"
	"os"
	"testing"

	"github.com/asahnoln/trello-updater/markdown"
	"github.com/asahnoln/trello-updater/trello"
)

var mdUrl string
var trelloUrl string
var boardId string

func init() {
	flag.StringVar(&mdUrl, "mdUrl", "", "markdown repository URL")
	flag.StringVar(&trelloUrl, "trelloUrl", "", "trello API URL")
	flag.StringVar(&boardId, "boardId", "", "trello board ID")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func checkFlags(t *testing.T) {
	switch {
	case mdUrl == "":
		t.Fatalf("-mdUrl flag should be set to markdown repository url")
	case trelloUrl == "":
		t.Fatalf("-trelloUrl flag should be set to API url")
	case boardId == "":
		t.Fatalf("-boardId flag should be set to trello board ID")
	}
}

func TestCreateData(t *testing.T) {
	mdReader := markdown.Parse(mdUrl)
	trello.Update(trelloUrl, boardId, mdReader)
}
