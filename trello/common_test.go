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
