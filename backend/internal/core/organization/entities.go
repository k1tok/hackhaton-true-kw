package organization

type Organization struct {
	ID             int            `json:"accountId"`
	IsCommercial   bool           `json:"isCommercial,omitempty"`
	Address        string         `json:"address"`
	BuildingType   string         `json:"buildingType,omitempty"`
	RoomsCount     string         `json:"roomsCount,omitempty"`
	ResidentsCount string         `json:"residentsCount,omitempty"`
	Consumption    map[string]int `json:"consumption,omitempty"`
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
