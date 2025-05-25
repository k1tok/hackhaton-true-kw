package organization

type ListOrgAPI interface {
	Search(address string, orgsChan chan<- []ListOrgAPIObj, errorChan chan<- error)
}

type OsmAPI interface {
	Search(address string, orgsChan chan<- []OsmObj, errorChan chan<- error)
}
