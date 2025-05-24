package osm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"true-kw/config"
	apperr "true-kw/errors"
	"true-kw/internal/core/organization"
)

// Структура для ответа Nominatim
type NominatimResult struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// Структуры для ответа Overpass
type OverpassResponse struct {
	Elements []struct {
		Type string            `json:"type"`
		ID   int64             `json:"id"`
		Tags map[string]string `json:"tags"`
		Lat  float64           `json:"lat"`
		Lon  float64           `json:"lon"`
	} `json:"elements"`
}

type OsmAPI struct {
	config     config.OsmAPI
	reqCounter int
}

func NewOsmAPI(config config.OsmAPI) *OsmAPI {

	return &OsmAPI{
		config: config,
	}
}
func (a *OsmAPI) Search(address string) ([]organization.OsmObj, error) {
	op := "OsmAPI.Search"
	if a.reqCounter == 6 {
		time.Sleep(1 * time.Second)
	}
	// 1. Получаем координаты через Nominatim
	nominatimURL := a.config.NominatimURL
	params := url.Values{}
	params.Add("q", address)
	params.Add("format", "json")
	params.Add("limit", "1")
	req, _ := http.NewRequest("GET", nominatimURL+"?"+params.Encode(), nil)
	req.Header.Set("User-Agent", "hakaton-true-kw/1.0 (anikocode@mail.ru)")

	client := &http.Client{}
	resp, err := client.Do(req)
	a.reqCounter++
	if err != nil {
		return nil, apperr.New(err, apperr.ErrInternal, "failed to make request", op)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var results []NominatimResult
	json.Unmarshal(body, &results)
	if len(results) == 0 {
		return nil, apperr.New(nil, apperr.ErrNotFound, "координаты не найдены", op)
	}
	lat, lon := results[0].Lat, results[0].Lon

	// 2. Формируем запрос к Overpass API
	// Ищем POI в радиусе 100 метров (можно увеличить)
	overpassQuery := `[out:json];
    (
      node(around:50,` + lat + `,` + lon + `)[amenity];
      node(around:50,` + lat + `,` + lon + `)[shop];
      node(around:50,` + lat + `,` + lon + `)[office];
    );
    out;`
	overpassURL := a.config.OverpassURL
	overpassResp, err := http.Post(overpassURL, "application/x-www-form-urlencoded", strings.NewReader("data="+url.QueryEscape(overpassQuery)))
	if err != nil {
		return nil, apperr.New(err, apperr.ErrInternal, "failed to make request", op)
	}
	defer overpassResp.Body.Close()
	overpassBody, _ := ioutil.ReadAll(overpassResp.Body)

	var overpassResult OverpassResponse
	json.Unmarshal(overpassBody, &overpassResult)

	// 3. Выводим организации
	if len(overpassResult.Elements) == 0 {
		return nil, apperr.New(nil, apperr.ErrNotFound, "организации не найдены", op)
	}
	var orgs []organization.OsmObj
	for _, elem := range overpassResult.Elements {
		name := elem.Tags["name"]
		if name == "" {
			name = "Без названия"
		}
		category := ""
		for _, key := range []string{"amenity", "shop", "office"} {
			if v, ok := elem.Tags[key]; ok {
				category = key + ": " + v
				break
			}
		}

		orgs = append(orgs, organization.OsmObj{Name: name, Type: category})
	}
	return orgs, nil
}
