package reqcrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetCompanyList() []string {
	var (
		domain       = "https://www.wantedly.com/"
		path         = "search?"
		page         = "1"
		query        = "page=" + page + "&q=python&t=projects"
		url          = domain + path + query
		companyList []string
	)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("not scrapping")
	}

	doc.Find(".company-name").Each(func(_ int, s *goquery.Selection) {
		if companyName := s.Text(); companyName != "" {
			companyName := strings.TrimRight(companyName, "\n")
			companyList = append(companyList, companyName)
		}
	})

	return companyList
}
