package organization

type Organization struct {
	ID             int            `json:"accountId"`
	IsCommercial   bool           `json:"isCommercial"`
	Adress         string         `json:"address"`
	BuildingType   string         `json:"buildingType"`
	RoomsCount     string         `json:"roomsCount"`
	ResidentsCount string         `json:"residentsCount"`
	Consumption    map[string]int `json:"consumption"`
}
