package organization

type Organization struct {
	ID             int            `json:"accountId"`
	IsCommercial   bool           `json:"isCommercial"`
	Address        string         `json:"address"`
	BuildingType   string         `json:"buildingType"`
	RoomsCount     string         `json:"roomsCount"`
	ResidentsCount string         `json:"residentsCount"`
	Consumption    map[string]int `json:"consumption"`
}

type OrganizationResult struct {
	ID           int     `json:"accountId"`
	IsCommercial bool    `json:"isCommercial"`
	Precision    float32 `json:"precision"`
}

type ListOrgAPIObj struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type OsmObj struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
