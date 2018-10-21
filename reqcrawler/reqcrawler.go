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
		company_list []string
	)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("not scrapping")
	}

	doc.Find(".company-name").Each(func(_ int, s *goquery.Selection) {
		if company_name := s.Text(); company_name != "" {
			company_name := strings.TrimRight(company_name, "\n")
			company_list = append(company_list, company_name)
		}
	})

	return company_list
}
