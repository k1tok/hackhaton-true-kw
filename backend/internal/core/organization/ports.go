package organization

type ListOrgAPI interface {
	Search(address string) ([]ListOrgAPIObj, error)
}

type OsmAPI interface {
	Search(address string) ([]OsmObj, error)
}
