package organization

import (
	"fmt"
	apperr "true-kw/errors"
)

type checkUseCase struct {
	listOrgAPI ListOrgAPI
	osmAPI     OsmAPI
}

func NewCheckUseCase(listOrgAPI ListOrgAPI, osmAPI OsmAPI) *checkUseCase {
	return &checkUseCase{
		listOrgAPI: listOrgAPI,
		osmAPI:     osmAPI,
	}
}

func (u *checkUseCase) Run(dta []Organization) ([]OrganizationResult, error) {
	op := "CheckUseCase.Run"
	if dta == nil {
		return nil, apperr.New(nil, apperr.ErrInternal, "data is empty", op)
	}
	osmOrgsChan := make(chan []OsmObj, 1)
	listOrgsChan := make(chan []ListOrgAPIObj, 1)
	errorOsmChan := make(chan error, 1)
	errorListOrgChan := make(chan error, 1)
	var result []OrganizationResult
	for _, org := range dta {
		if org.Address == "" {
			continue
		}
		go u.osmAPI.Search(org.Address, osmOrgsChan, errorOsmChan)
		go u.listOrgAPI.Search(org.Address, listOrgsChan, errorListOrgChan)

		if err := <-errorOsmChan; err != nil {
			if !apperr.IsErrorCode(err, apperr.ErrNotFound) {
				return nil, err
			}
		}
		if err := <-errorListOrgChan; err != nil {
			if !apperr.IsErrorCode(err, apperr.ErrNotFound) {
				return nil, err
			}
		}
		osmOrgs := <-osmOrgsChan
		listOrgs := <-listOrgsChan

		fmt.Println(len(listOrgs))
		fmt.Println(len(osmOrgs))
		result = append(result, OrganizationResult{org.ID, false, 0.0})
	}
	return result, nil
}
