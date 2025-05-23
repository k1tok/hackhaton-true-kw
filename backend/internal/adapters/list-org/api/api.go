package listorg

import (
	"fmt"
	"net/http"
	"net/url"
	"true-kw/config"
	apperr "true-kw/errors"
	"true-kw/internal/core/organization"

	"github.com/PuerkitoBio/goquery"
)

type ListOrgAPI struct {
	config config.ListOrgAPI
}

func NewListOrgAPI(config config.ListOrgAPI) *ListOrgAPI {
	return &ListOrgAPI{
		config: config,
	}
}

func (a *ListOrgAPI) Search(address string) ([]organization.ListOrgAPIObj, error) {
	op := "ListOrgAPI.Search"
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
	if err != nil {
		return nil, apperr.New(err, apperr.ErrInternal, "failed to make request", op)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, apperr.New(err, apperr.ErrInternal, "failed to parse HTML", op)
	}
	var orgs []organization.ListOrgAPIObj

	doc.Find("div.org_list a").Each(func(i int, s *goquery.Selection) {
		orgName := s.Text()
		orgLink, _ := s.Attr("href")
		orgs = append(orgs, organization.ListOrgAPIObj{Name: orgName, Link: orgLink})
		fmt.Printf("Организация: %s (%s)\n", orgName, orgLink)
	})
	return orgs, nil
}
