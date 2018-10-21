package reqcrawler

import (
	"testing"
	"github.com/tarunama/reqcrawler/reqcrawler"
)

func TestGetCompanyList(t *testing.T) {
	actual := reqcrawler.GetCompanyList()
	expected := 1

	if expected <= len(actual) {
		t.Errorf("No get company in GetCompanyList %s", actual)
	}
}
