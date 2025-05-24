package listorg

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
	"true-kw/config"
	apperr "true-kw/errors"
	"true-kw/internal/core/organization"

	"github.com/PuerkitoBio/goquery"
)

type ListOrgAPI struct {
	config     config.ListOrgAPI
	regex      *regexp.Regexp
	reqCounter int
}

func NewListOrgAPI(config config.ListOrgAPI) *ListOrgAPI {
	re := regexp.MustCompile(`[Дд]\. (\d+)`)
	return &ListOrgAPI{
		config: config,
		regex:  re,
	}
}

func (a *ListOrgAPI) Search(address string) ([]organization.ListOrgAPIObj, error) {
	op := "ListOrgAPI.Search"
	if a.reqCounter == 10 {
		time.Sleep(1 * time.Second)
	}
	baseURL := a.config.URL
	params := url.Values{}
	params.Add("type", "address")
	params.Add("val", address)
	searchURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, apperr.New(err, apperr.ErrInternal, "failed to create request", op)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; hakaton-true-kw/1.0; +http://your-site.ru/)")

	client := &http.Client{}
	resp, err := client.Do(req)
	a.reqCounter++
	if err != nil {
		return nil, apperr.New(err, apperr.ErrInternal, "failed to make request", op)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, apperr.New(err, apperr.ErrInternal, "failed to parse HTML", op)
	}
	var orgs []organization.ListOrgAPIObj

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		a_tag := s.Find("a")
		orgName := a_tag.Text()
		orgLink, _ := a_tag.Attr("href")
		legalAddress := ""
		s.Find("span").Each(func(i int, s *goquery.Selection) {
			if strings.Contains(s.Text(), "юр.адрес") {
				// Разбиваем текст по строкам (разделитель <br>)
				textLines := strings.Split(s.Text(), "\n")
				for _, line := range textLines {
					if strings.Contains(line, "юр.адрес") {
						// Извлекаем часть после "юр.адрес:"
						parts := strings.SplitN(line, "юр.адрес:", 2)
						if len(parts) > 1 {
							legalAddress = strings.TrimSpace(parts[1])
							break
						}
					}
				}
			}
		})
		// Извлекаем номер дома из адреса
		matches := a.regex.FindStringSubmatch(legalAddress)
		if len(matches) >= 2 {
			houseNumber := matches[1]
			matches = a.regex.FindStringSubmatch(address)
			if len(matches) >= 2 {
				houseNumberInAddress := matches[1]
				if houseNumber == houseNumberInAddress {
					orgs = append(orgs, organization.ListOrgAPIObj{Name: orgName, Link: orgLink})
				}
			}
		}
	})
	return orgs, nil
}
